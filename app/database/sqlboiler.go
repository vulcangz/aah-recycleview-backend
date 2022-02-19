package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	aah "aahframe.work"
	"github.com/volatiletech/sqlboiler/boil"
)

// SQLBoiler public declaration
var DB *sql.DB

// New method creates database connection with given aah application instance :)
func SConnect(_ *aah.Event) {
	cfg := aah.App().Config()

	var (
		dsn string
		err error
	)

	db_connection := cfg.StringDefault("database.driver", "mysql")

	switch db_connection {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.StringDefault("database.username", ""),
			cfg.StringDefault("database.password", ""),
			cfg.StringDefault("database.host", "localhost"),
			cfg.StringDefault("database.port", "3306"),
			cfg.StringDefault("database.name", ""),
		)
		DB, err = sql.Open("mysql", dsn)
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			cfg.StringDefault("database.host", "localhost"),
			cfg.StringDefault("database.username", ""),
			cfg.StringDefault("database.password", ""),
			cfg.StringDefault("database.name", ""),
			cfg.StringDefault("database.port", "3306"),
		)
		DB, err = sql.Open("postgres", dsn)
	case "sqlite":
		DB, err = sql.Open("sqlite", fmt.Sprintf("%v", cfg.StringDefault("database.name", "test.db")))
	default:
		aah.App().Log().Error("Fatal error database connetion: not supported database adapter")
		panic(errors.New("not supported database adapter"))
	}

	if err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}

	if err := DB.Ping(); err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}

	// Just for debugging
	if err == nil {
		aah.App().Log().Debug("Database connection successful!")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	DB.SetMaxIdleConns(cfg.IntDefault("database.max_idle_connections", 2))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	DB.SetMaxOpenConns(cfg.IntDefault("database.max_active_connections", 10))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	d := time.Duration(cfg.IntDefault("database.max_connection_lifetime", 0))
	DB.SetConnMaxLifetime(d * time.Hour)

	boil.SetDB(DB)

}

// GetDB - get a connection
func SDB() *sql.DB {
	return DB
}

// Close the database
func SDisconnect(_ *aah.Event) {
	DB.Close()
	aah.App().Log().Debug("Database closed successful!")
}
