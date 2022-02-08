package database

import (
	"errors"
	"fmt"
	"time"

	aah "aahframe.work"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Gorm public declaration
var Gorm *gorm.DB

// New method creates database connection with given aah application instance :)
func Connect(_ *aah.Event) {
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
		Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			cfg.StringDefault("database.host", "localhost"),
			cfg.StringDefault("database.username", ""),
			cfg.StringDefault("database.password", ""),
			cfg.StringDefault("database.name", ""),
			cfg.StringDefault("database.port", "3306"),
		)
		Gorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		Gorm, err = gorm.Open(sqlite.Open(fmt.Sprintf("%v", cfg.StringDefault("database.name", "eav.db"))), &gorm.Config{})
	default:
		aah.App().Log().Error("Fatal error database connetion: not supported database adapter")
		panic(errors.New("not supported database adapter"))
	}

	if err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}

	sqlDB, err := Gorm.DB()
	if err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}

	// Just for debugging
	if err == nil {
		aah.App().Log().Debug("Database connection successful!")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.IntDefault("database.max_idle_connections", 2))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.IntDefault("database.max_active_connections", 10))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	d := time.Duration(cfg.IntDefault("database.max_connection_lifetime", 0))
	sqlDB.SetConnMaxLifetime(d * time.Hour)
}

// GetDB - get a connection
func GetDB() *gorm.DB {
	return Gorm
}

// Close the database
func Disconnect(_ *aah.Event) {
	sqlDB, err := GetDB().DB()
	if err != nil {
		aah.App().Log().Errorf("Fatal error database connetion: %s", err.Error())
	}
	sqlDB.Close()
	aah.App().Log().Debug("Database closed successful!")
}
