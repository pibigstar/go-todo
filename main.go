package main

import (
	"gitee.com/johng/gf/g"
	"github.com/pibigstar/go-todo/config"
	_ "github.com/pibigstar/go-todo/controller"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("main")

func main() {
	s := g.Server()
	port := config.ServerConfig.Port
	s.SetPort(int(port))
	host := config.ServerConfig.Host
	s.Domain(host)

	// 开启日志
	s.SetLogPath("log/todo.log")
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)

	s.Run()

}
