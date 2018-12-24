package test

import (
	"testing"

	"gitee.com/johng/gf/g"
	"github.com/pibigstar/go-todo/config"
)

func TestMain(t *testing.T) {
	s := g.Server()
	port := config.ServerConfig.Port

	s.SetPort(int(port))
	s.Run()
}
