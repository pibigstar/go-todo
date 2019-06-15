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
	s.BindHandler("/api/user/delete", adminDelete)
	s.BindHandler("/api/user/blacklist", blackList)
}

type AdminLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type IDRequest struct {
	ID int `json:"id"`
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

func adminDelete(r *ghttp.Request) {
	request := new(IDRequest)
	r.GetJson().ToStruct(request)
	if request.ID == 0 {
		return
	}
	err := models.MAdmin.AdminDelete(request.ID)
	if err != nil {
		log.Error("delete user failed")
	}
	utils.SuccessResponse("OK")
}

