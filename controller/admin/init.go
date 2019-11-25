package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("admin")

func init() {
	s := g.Server()
	s.BindHookHandler("/api/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		// 跨域访问设置
		r.Response.CORSDefault()
		r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	})
}
