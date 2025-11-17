package client

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2" 
)

var (
	chOnce   sync.Once
	chClient ClickHouse
)

type ClickHouse interface {
	Initialize(ctx context.Context) error
	Ping(ctx context.Context) error
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	WithTx(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error
	InsertBatch(ctx context.Context, table string, columns []string, rows [][]any) error
	Close() error
}


type clickhouseImpl struct {
	db *sql.DB
}

func GetClickHouseInstance() ClickHouse {
	chOnce.Do(func() {
		chClient = &clickhouseImpl{}
	})
	return chClient
}


func (c *clickhouseImpl) Initialize(ctx context.Context) error {
	dsn := os.Getenv("CLICKHOUSE_DSN")
	if dsn == "" {
		dsn = "clickhouse://otel:otelpwd@clickhouse:9000/otel?dial_timeout=3s&compress=true"
	}

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return fmt.Errorf("clickhouse open: %w", err)
	}
	// pool tuning (ajuste conforme tráfego)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	// tenta ping
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctxPing); err != nil {
		_ = db.Close()
		return fmt.Errorf("clickhouse ping: %w", err)
	}

	c.db = db
	return nil
}

func (c *clickhouseImpl) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *clickhouseImpl) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *clickhouseImpl) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}

func (c *clickhouseImpl) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}


func (c *clickhouseImpl) WithTx(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}


func (c *clickhouseImpl) InsertBatch(ctx context.Context, table string, columns []string, rows [][]any) error {
	if len(rows) == 0 {
		return nil
	}
	
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		joinColumns(columns),
		joinColumns(placeholders),
	)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare batch: %w", err)
	}
	defer stmt.Close()

	for _, row := range rows {
		if _, err := stmt.ExecContext(ctx, row...); err != nil {
			return fmt.Errorf("batch exec row: %w", err)
		}
	}

	return nil
}

func (c *clickhouseImpl) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}


func joinColumns(xs []string) string {
	if len(xs) == 0 {
		return ""
	}
	out := xs[0]
	for i := 1; i < len(xs); i++ {
		out += "," + xs[i]
	}
	return out
}
