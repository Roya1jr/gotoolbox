package gtbdb

import "time"

type Dsnconfig struct {
	DBName   string
	Username string
	Password string
	Hostname string
	Protocol string
}

type Databaseclientconfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	DriverName      string
}
