package test

import (
	"fmt"
	"testing"

	"github.com/pibigstar/go-todo/config"
)

func TestConfig(t *testing.T) {
	fmt.Printf("%+v", config.ServerConfig)
}
