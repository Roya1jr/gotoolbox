package gtb_db

import "time"

type Dsnconfig struct {
	DbName   string
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
