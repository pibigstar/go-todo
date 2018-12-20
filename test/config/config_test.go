package test

import (
	"testing"

	"github.com/pibigstar/go-todo/config"
)

func TestConfig(t *testing.T) {
	mysqlCfg := config.LoadConfig()
	log.Info("db config", db)
}
