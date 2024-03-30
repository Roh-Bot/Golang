package messagebrokers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Roh-bot/Golang/config"
	"github.com/Roh-bot/Golang/db"
	"github.com/Roh-bot/Golang/db/pgfuncs"
	"github.com/Roh-bot/Golang/loggers"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	RabbitMqPublisher *publisher
	RabbitMq          *rabbitMq
)

type rabbitMq struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
	p    *publisher
}

type publisher struct {
	messages chan map[string]any
	subs     []chan map[string]any
	mutex    *sync.RWMutex
}

func NewRabbitInstance() error {
	log.Println("Connecting to RabbitMq")
	conn, err := amqp091.Dial(config.RABBITURL)
	if err != nil {
		loggers.Zap.Errorf("RabbitMQ Error: %s", err.Error())
		return err
	}

	if conn.IsClosed() {
		loggers.Zap.Errorf("RabbitMQ connection is closed")
		return fmt.Errorf("RabbitMQ connection is closed")
	}

	ch, err := conn.Channel()
	if err != nil {
		loggers.Zap.Errorf("RabbitMQ Error: %s", err.Error())
		return err
	}

	RabbitMq = &rabbitMq{
		conn: conn,
		ch:   ch,
	}

	log.Println("Connected to RabbitMQ")
	return nil
}

func (r *rabbitMq) NewPublisher() {
	r.p = &publisher{
		messages: make(chan map[string]any),
		subs:     make([]chan map[string]any, 0),
		mutex:    new(sync.RWMutex),
	}
	RabbitMqPublisher = r.p
}

func (r *rabbitMq) SendMessages(body map[string]any) {
	if r.conn.IsClosed() || r.ch.IsClosed() {
		loggers.Zap.Errorf("RabbitMQ connection/channel is closed")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return
	}

	var queues []string
	rows, err := db.PostgresPool.Read(pgfuncs.GetAllQueues)
	if err != nil {
		loggers.Zap.Errorf("PG Error: %s", err.Error())
		return
	}
	for rows.Next() {
		err := rows.Scan(&queues)
		if err != nil {
			loggers.Zap.Errorf("PG Error: %s", err.Error())
			return
		}
	}

	for _, queue := range queues {
		err := r.ch.QueueBind(
			queue,
			"",
			config.RABBITEXCHANGE,
			false,
			nil,
		)
		if err != nil {
			return
		}
	}

	err = r.ch.PublishWithContext(
		ctx,
		config.RABBITEXCHANGE,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		})
	if err != nil {
		loggers.Zap.Errorf("%s : %s", err, "Failed to publish a message")
		return
	}

	log.Printf("RABBITMQ ORDERS SENT: Sent %v\n", body)
	return
}

func (r *rabbitMq) ReceiveMessages() {
	if r.conn.IsClosed() || r.ch.IsClosed() {
		log.Println("RabbitMQ connection is already closed")
		return
	}

	var queueName string
	err := db.PGReadSingleRow(&queueName, pgfuncs.GetQueueName)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err = db.PostgresPool.CreateUpdateDelete(pgfuncs.DeleteQueue, queueName)
		if err != nil {
			loggers.Zap.Errorf("PG Error: %s", err.Error())
			return
		}
	}()

	q, err := r.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,
	)
	if err != nil {
		return
	}
	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,
	)
	if err != nil {
		loggers.Zap.Errorf("RabbitMQ Error: %s", err.Error())
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	go func() {
		for deliver := range msgs {
			body := make(map[string]any)
			if err := json.Unmarshal(deliver.Body, &body); err != nil {
				loggers.Zap.Errorf("RabbitMQ Error: %s", err.Error())
				continue
			}
			r.p.publishMessage(body)
			log.Printf("Received a message: %v", body)
		}
		log.Println("Exiting RabbitMQ receiver...")
	}()

	<-quit
	log.Println("Closing RabbitMQ receiver...")
}

func (r *rabbitMq) CloseConnection() {
	log.Println("Closing rabbit connection...")
	if r.conn.IsClosed() {
		log.Println("RabbitMQ connection is already closed")
		return
	}
	err := r.conn.Close()
	if err != nil {
		loggers.Zap.Errorf("Failed to close connection: %s", err)
	}
	log.Println("Rabbit connection closed successfully")

	log.Println("Closing rabbit channel...")
	err = r.ch.Close()
	if err != nil {
		loggers.Zap.Errorf("Failed to close channel: %s", err)
	}
	log.Println("Rabbit channel closed successfully")
}

func (p *publisher) publishMessage(body map[string]any) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	for _, sub := range p.subs {
		sub <- body
	}
}

func (p *publisher) SubscribeMessages(sub chan map[string]any) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.subs = append(p.subs, sub)
}

func (p *publisher) Shutdown() {
	log.Println("Closing all RabbitMq messaging subscribers")
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, sub := range p.subs {
		close(sub)
	}
	log.Println("Closed all RabbitMq messaging subscribers")
}
