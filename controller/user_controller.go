package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
	"gitee.com/johng/gf/g/util/gvalid"
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/wxLogin", wxLogin)
}

// WxLoginRequest 微信登录request
type WxLoginRequest struct {
	Code string `json:"code" gvalid:"type@required#code码不能为空"`
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

// WxLogin 微信登录
func wxLogin(r *ghttp.Request) {

	wxLoginRequest := new(WxLoginRequest)
	r.GetJson().ToStruct(wxLoginRequest)

	if err := gvalid.CheckStruct(wxLoginRequest, nil); err != nil {
		log.Error("code为空", "err", err.String())
		r.Response.WriteJson(utils.ErrorResponse(err.String()))
		return
	}
	var wxLoginResp WxLoginResponse
	// 拿到session_key 和 openid
	client := &http.Client{}
	url := fmt.Sprintf(config.ServerConfig.WxLoginURL, config.ServerConfig.Appid, config.ServerConfig.Secret, wxLoginRequest.Code)
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
	wxLoginResp.Token = token
	r.Response.WriteJson(wxLoginResp)
}
