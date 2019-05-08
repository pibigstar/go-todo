package admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/api/login", adminLogin)
	s.BindHandler("/api/user/list", adminList)
	s.BindHandler("/api/user/blacklist", blackList)
}

type AdminLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func adminLogin(r *ghttp.Request) {
	request := new(AdminLoginRequest)
	r.GetToStruct(request)
	admin, err := models.MAdmin.Login(request.UserName, request.Password)
	if err != nil {
		log.Error("login is failed", "err", err.Error())
		utils.Error(r)
	}
	r.Response.WriteJson(utils.SuccessWithData("OK", admin))
}

func adminList(r *ghttp.Request) {

	admins, err := models.MAdmin.ListAdmin()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, admins)
}

func blackList(r *ghttp.Request) {
	blacks, err := models.MBlackUser.ListBlack()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, blacks)
}
