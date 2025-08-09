package gtb_db

import (
	"database/sql"
	"fmt"
)

// CreateDsn securely builds the Data Source Name.
func (cfg *Dsnconfig) CreateDsn() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s",
		cfg.Username, cfg.Password, cfg.Protocol, cfg.Hostname, cfg.DbName)
}

// GetDatabaseClient opens and configures a database client with
func GetDatabaseClient(dsn string, cfg *Databaseclientconfig) (*sql.DB, error) {
	client, err := sql.Open(cfg.DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %w", err)
	}

	client.SetMaxOpenConns(cfg.MaxOpenConns)
	client.SetMaxIdleConns(cfg.MaxIdleConns)
	client.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := client.Ping(); err != nil {
		client.Close()
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return client, nil
}
