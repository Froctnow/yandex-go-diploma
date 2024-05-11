package pgclient

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"io/fs"
)

type Client struct {
	*sqlx.DB
	parser Parser
}

type PostgreSQL struct {
	ConnString     string
	PathsToQueries []string

	LogLevel LogLevel
}

// PGClient wrapped sqlx and pgx clients
type PGClient interface {
	NamedExec(
		ctx context.Context,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args interface{},
	) (sql.Result, error)
	Exec(
		ctx context.Context,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args ...any,
	) (sql.Result, error)
	NamedQuery(
		ctx context.Context,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args interface{},
	) (*sqlx.Rows, error)
	NamedQueryxContext(
		ctx context.Context,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args ...interface{},
	) (*sqlx.Rows, error)
	NamedGetContext(
		ctx context.Context,
		dest any,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args ...interface{},
	) error
	NamedSelectContext(
		ctx context.Context,
		dest any,
		queryName string,
		templateParams map[string]interface{},
		transaction Transaction,
		args ...interface{},
	) error

	BeginTransaction() (Transaction, error)
	CloseConnections() error
	GetQueryByName(
		name string,
		params map[string]interface{},
	) (string, error)
	HealthCheckConnection() error
}

// Transaction wrap sqlx.Tx
type Transaction interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Commit() error
	Rollback() error
}

const (
	driverName = "pgx"

	SQLFileExt = "*.sql"
)

func New(
	cfg PostgreSQL,
	log LogClient,
	queryFiles fs.FS,
) (PGClient, error) {
	if len(cfg.PathsToQueries) == 0 {
		return nil, fmt.Errorf("empty param PathsToQueries")
	}

	parser := NewParser()
	if queryFiles == nil {
		for _, path := range cfg.PathsToQueries {
			err := parser.AddRoot(path, SQLFileExt)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := parser.AddFSRoot(cfg.PathsToQueries, queryFiles, SQLFileExt)
		if err != nil {
			return nil, err
		}
	}

	connConfig, err := pgx.ParseConfig(cfg.ConnString)
	if err != nil {
		return nil, err
	}

	if log != nil {
		adapter := NewLogAdapter(log)
		connConfig.Logger = adapter
		connConfig.LogLevel = pgx.LogLevel(cfg.LogLevel)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, err
	}

	sqlxDB := sqlx.NewDb(db, driverName)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Client{
		sqlxDB,
		parser,
	}, nil
}

// NamedExec uses sqlx.NamedExecContext
func (c *Client) NamedExec(
	ctx context.Context,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args interface{},
) (sql.Result, error) {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return nil, err
	}
	if query == "" {
		return nil, fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.NamedExecContext(ctx, query, args)
	}
	return c.NamedExecContext(ctx, query, args)
}

// Exec uses sqlx.ExecContext
func (c *Client) Exec(
	ctx context.Context,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args ...any,
) (sql.Result, error) {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return nil, err
	}
	if query == "" {
		return nil, fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.ExecContext(ctx, query, args...)
	}

	return c.ExecContext(ctx, query, args...)
}

// NamedQuery uses sqlx.NamedQueryContext and NamedQuery for transactions
func (c *Client) NamedQuery(
	ctx context.Context,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args interface{},
) (*sqlx.Rows, error) {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return nil, err
	}
	if query == "" {
		return nil, fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.NamedQuery(query, args)
	}
	return c.NamedQueryContext(ctx, query, args)
}

// NamedQueryxContext uses sqlx.QueryxContext
func (c *Client) NamedQueryxContext(
	ctx context.Context,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args ...interface{},
) (*sqlx.Rows, error) {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return nil, err
	}
	if query == "" {
		return nil, fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.QueryxContext(ctx, query, args...)
	}

	return c.QueryxContext(ctx, query, args...)
}

// NamedGetContext uses sqlx.GetContext
func (c *Client) NamedGetContext(
	ctx context.Context,
	dest any,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args ...interface{},
) error {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return err
	}
	if query == "" {
		return fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.GetContext(ctx, dest, query, args...)
	}

	return c.GetContext(ctx, dest, query, args...)
}

// NamedSelectContext uses sqlx.SelectContext
func (c *Client) NamedSelectContext(
	ctx context.Context,
	dest any,
	queryName string,
	templateParams map[string]interface{},
	transaction Transaction,
	args ...interface{},
) error {
	query, err := c.parser.Exec(queryName, templateParams)
	if err != nil {
		return err
	}
	if query == "" {
		return fmt.Errorf("not found query by name: %v", queryName)
	}

	if transaction != nil {
		return transaction.SelectContext(ctx, dest, query, args...)
	}

	return c.SelectContext(ctx, dest, query, args...)
}

// BeginTransaction create transaction *sqlx.TX
func (c *Client) BeginTransaction() (Transaction, error) {
	return c.Beginx()
}

// CloseConnections run (*sql.DB).Close()
func (c *Client) CloseConnections() error {
	return c.Close()
}

// GetQueryByName return parsed query text from templates
func (c *Client) GetQueryByName(
	name string,
	params map[string]interface{},
) (string, error) {
	return c.parser.Exec(name, params)
}

// HealthCheckConnection return error if connection is closed
func (c *Client) HealthCheckConnection() error {
	err := c.Ping()

	if err == nil {
		return nil
	}

	return fmt.Errorf("connection to database is broken, err %w", err)
}
