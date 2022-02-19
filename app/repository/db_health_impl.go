package repository

import (
	"context"
	"errors"
	"strconv"

	aah "aahframe.work"
	"github.com/mediocregopher/radix/v4"
)

const (
	// Scrape query.
	globalStatusQuery = `SHOW GLOBAL STATUS`
)

func (repo *SormRepository) RdbHealthCheck() (*Health, error) {
	var (
		h            *Health
		queryCount   int
		slaveRunning bool
	)
	slaveRunning = false

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB := repo.db

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// https://github.com/fuze/mysql-healthcheck/blob/master/main.go
	globalStats, err := sqlDB.Query(globalStatusQuery)
	if err != nil {
		aah.App().Log().Debugf("MySQL Running: [false], Query Count: [0],  Slave Running: [0]")
	}
	defer globalStats.Close()

	for globalStats.Next() {
		var name string
		var value string
		err = globalStats.Scan(&name, &value)

		// Check current queries
		if name == "Threads_connected" {
			queryCount, err = strconv.Atoi(value)
		}
		// Check if slave running
		if name == "Slave_running" && value == "ON" {
			slaveRunning = true
		}

		if err != nil {
			aah.App().Log().Debugf("%v", err)
			h = &Health{
				Status:       "bad",
				QueryCount:   -1,
				SlaveRunning: false,
			}
			return h, nil
		}
	}
	aah.App().Log().Debugf("MySQL Running: [true], Query Count: [%v],  Slave Running: [%v]", queryCount, slaveRunning)

	h = &Health{
		Status:       "ok",
		QueryCount:   queryCount,
		SlaveRunning: slaveRunning,
	}
	return h, nil
}

func RedisGetSet(ctx context.Context, client radix.Client, key, val string) error {
	if err := client.Do(ctx, radix.Cmd(nil, "SET", key, val)); err != nil {
		return err
	}
	var out string
	if err := client.Do(ctx, radix.Cmd(&out, "GET", key)); err != nil {
		return err
	} else if out != val {
		return errors.New("got wrong value")
	}
	return nil
}
