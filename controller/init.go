package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("controller")

func init() {
	s := g.Server()
	s.BindHookHandler("/send", ghttp.HOOK_BEFORE_SERVE, middleware.CheckToken)
	s.BindHookHandler("/task/*", ghttp.HOOK_BEFORE_SERVE, middleware.CheckToken)
	s.BindHookHandler("/group/*", ghttp.HOOK_BEFORE_SERVE, middleware.CheckToken)
}
