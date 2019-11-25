package test

import (
	"testing"

	"github.com/pibigstar/go-todo/config"
)

func TestConfig(t *testing.T) {
	t.Logf("%+v", config.ServerConfig)
}
