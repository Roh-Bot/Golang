package config

// PostgreSQL configuration
const (
	PostgreSQLHost     = "localhost"
	PostgreSQLDatabase = "postgres"
	PostgreSQLUser     = "postgres"
	PostgreSQLPassword = "admin"
)

// MESSAGE BROKERS
const (
	TopicKafka          = "Notifications"
	RABBITURL           = "amqp://guest:guest@127.0.0.1:5672/"
	RABBITQUEUE         = "BrokerOrdersBeta"
	RABBITEXCHANGE      = "Publisher"
	RABBITORDERSNOTIKEY = "brokers.orders.noti.beta"
	KAFKAURL            = "localhost:9092"
)
