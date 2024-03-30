package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/Roh-bot/Golang/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

var (
	PostgresPool *postgresPoolClient
)

type postgresPoolClient struct {
	pool *pgxpool.Pool
}

// postgresNewPool creates a connection for PostgreSQL.
// Make sure to close the connection when you no longer require it.
// To close a connection use CloseConnection method.
func postgresNewPool() error {
	configString := fmt.Sprintf(`host=%s  user=%s password=%s dbname=%s sslmode=disable`,
		config.PostgreSQLHost, config.PostgreSQLUser, config.PostgreSQLPassword, config.PostgreSQLDatabase)

	configDb, err := pgxpool.ParseConfig(configString)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, configDb)
	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	err = pool.Ping(ctx)
	if err != nil {
		return err
	}

	PostgresPool = &postgresPoolClient{pool: pool}
	return nil
}

// Read is intended for performing read operations on the database.
// Use this function only when connection pooling is required.
func (p *postgresPoolClient) Read(call string, data ...any) (pgx.Rows, error) {
	if p.pool == nil {
		return nil, errors.New("no connections available in the pool")
	}
	ctx := context.Background()
	rows, err := p.pool.Query(ctx, call, data...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// CreateUpdateDelete is intended for performing create, update and delete functions.
//
// ------------------ IMPORTANT ----------------
//
//	DO NOT USE THIS FUNCTION TO PERFORM
//	READ OPERATIONS. DOING SO WILL
//	RESULT IN UNEXPECTED RESULTS.
//
// ------------------ IMPORTANT ----------------
func (p *postgresPoolClient) CreateUpdateDelete(call string, data ...any) error {
	if p.pool == nil {
		return errors.New("no connections available in the pool")
	}
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	_, err := p.pool.Exec(ctx, call, data...)
	if err != nil {
		return err
	}

	//log.Print(cmdTag.RowsAffected())

	return nil
}

// PGReadSingleRow is intended for performing a read operation
// on a query which strictly returns scalar data.
func PGReadSingleRow[T comparable](scanData *T, call string, data ...any) error {
	if PostgresPool.pool == nil {
		return errors.New("no connections available in the pool")
	}
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	row := PostgresPool.pool.QueryRow(ctx, call, data...)

	err := row.Scan(scanData)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresPoolClient) FlushPool() {
	if p.pool == nil {
		slog.Error("No connection to close")
		return
	}
	p.pool.Close()
	return
}
