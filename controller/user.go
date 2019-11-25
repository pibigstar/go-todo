package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/wxLogin", wxLogin)
	s.BindHandler("/phoneLogin", phoneLogin)
	s.BindHandler("/user/info", getUserInfo)
	s.BindHandler("/user/update", updateUserInfo)
}

// WxLoginRequest 微信登录request
type WxLoginRequest struct {
	Code      string `json:"code" gvalid:"type@required#code码不能为空"`
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
}

// WxLoginResponse 微信登录response
type WxLoginResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	ErrMsg     string `json:"errMsg"`
	Token      string `json:"token"`
}

type PhoneLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUserInfoRequest struct {
	NickName      string `json:"nickName"`
	RealName      string `json:"realName"`
	Gender        int    `json:"gender"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	ReceiveRemind bool   `json:"receiveRemind"`
}

// WxLogin 微信登录
func wxLogin(r *ghttp.Request) {

	wxLoginRequest := &WxLoginRequest{}
	r.GetToStruct(wxLoginRequest)

	if err := gvalid.CheckStruct(wxLoginRequest, nil); err != nil {
		log.Error("code为空", "err", err.String())
		r.Response.WriteJson(utils.ErrorResponse(err.String()))
		return
	}
	var wxLoginResp WxLoginResponse
	// 拿到session_key 和 openid
	client := &http.Client{}
	url := fmt.Sprintf(config.ServerConfig.WxLoginURL, config.ServerConfig.Appid, config.ServerConfig.Secret, wxLoginRequest.Code)
	fmt.Println(url)
	res, err := client.Get(url)
	if err != nil {
		log.Error("获取openId失败", "err", err.Error())
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &wxLoginResp)
	token, err := utils.GenOpenIDToken(wxLoginResp.Openid)
	if err != nil {
		log.Error("生成token失败", "err", err.Error())
	}
	//判断用户是否是新用户
	_, err = models.MUser.GetUserByOpenID(wxLoginResp.Openid)
	if err != nil && gorm.IsRecordNotFoundError(err) {
		//将此用户插入到数据库中
		newUser := &models.User{
			OpenID:     wxLoginResp.Openid,
			NickName:   wxLoginRequest.NickName,
			AvatarURL:  wxLoginRequest.AvatarURL,
			Gender:     wxLoginRequest.Gender,
			UpdateTime: time.Now(),
			CreateTime: time.Now(),
		}
		err = models.MUser.Create(newUser)
		if err != nil {
			log.Error("insert user failed", "err", err.Error())
		}
	}
	wxLoginResp.Token = token
	r.Response.WriteJson(wxLoginResp)
}

func phoneLogin(r *ghttp.Request) {
	phoneLoginRequest := new(PhoneLoginRequest)
	r.GetToStruct(phoneLoginRequest)

	user, err := models.MUser.PhoneLogin(phoneLoginRequest.Phone, phoneLoginRequest.Password)
	if err != nil {
		log.Error("user login is failed", "phone", phoneLoginRequest.Phone, "password", phoneLoginRequest.Password)
		r.Response.WriteJson(utils.ErrorResponse("账号或密码错误"))
		r.Exit()
	} else {
		token, err := utils.GenOpenIDToken(user.OpenID)
		if err != nil {
			log.Error("gender token is failed")
		}
		r.Response.WriteJson(utils.SuccessWithData("OK", token))
	}
}

func getUserInfo(r *ghttp.Request) {
	openId, err := middleware.GetOpenID(r)
	if err != nil {
		log.Error("get user openId is failed", "err", err.Error())
		r.Exit()
	}
	user, err := models.MUser.GetUserByOpenID(openId)
	if err != nil {
		log.Error("get user info is failed", "openId", openId)
	}
	r.Response.WriteJson(utils.SuccessWithData("OK", user))
}

func updateUserInfo(r *ghttp.Request) {
	updateUserInfoRequest := new(UpdateUserInfoRequest)
	r.GetToStruct(updateUserInfoRequest)

	openId, err := middleware.GetOpenID(r)
	if err != nil {
		log.Error("get user openId is failed", "err", err.Error())
		r.Exit()
	}
	model := convertRequestToModel(updateUserInfoRequest)
	model.OpenID = openId
	err = models.MUser.UpdateUserInfo(model)
	if err != nil {
		log.Error("update user info is failed", "openId", openId)
	}
	r.Response.WriteJson(utils.SuccessResponse("OK"))
}
func convertRequestToModel(request *UpdateUserInfoRequest) *models.User {
	return &models.User{
		NickName: request.NickName,
		RealName: request.RealName,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Password: request.Password,
	}
}
