package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// WxLogin 微信登录
func WxLogin(ctx *gin.Context) {

	type WxLogin struct {
		Code string `json:"code" binding:"required"`
	}
	type WxLoginResponse struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		Unionid    string `json:"unionid"`
		Errcode    int    `json:"errcode"`
		ErrMsg     string `json:"errMsg"`
	}
	//var wxLogin WxLogin
	var wxLoginResp WxLoginResponse

	code := ctx.Param("code")

	appID := viper.GetStringMap("common")["appid"]
	secret := viper.GetStringMap("common")["secret"]

	fmt.Printf("appId:%s, secret:%s", appID, secret)
	// 拿到session_key 和 openid
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, secret, code)
	res, err := client.Get(url)
	if err != nil {
		log.CtxError(ctx, "获取openId失败", "err", err.Error())
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &wxLoginResp)

	fmt.Printf("%+v", wxLoginResp)

}
