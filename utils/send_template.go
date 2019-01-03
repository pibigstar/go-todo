package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/constant"
	"github.com/pibigstar/go-todo/models/db"
)

// 发送模板消息

var (
	send_template_url    = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s"
	get_access_token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// SendTemplate 发送模板消息
func SendTemplate(msg *TemplateMsg) (*SendTemplateResponse, error) {
	msg.Miniprogram.AppID = config.ServerConfig.Appid
	accessToken, err := getAccessToken(msg.Touser)
	if err != nil {
		log.Error("获取accessToken失败")
		return nil, err
	}
	url := fmt.Sprintf(send_template_url, accessToken.AccessToken)
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("模板消息JSON格式错误", "err", err.Error())
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Error("网络错误，发送模板消息失败", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var templateResponse SendTemplateResponse
	err = json.Unmarshal(body, &templateResponse)
	if err != nil {
		log.Error("解析responseBody错误", "err", err.Error())
		return nil, err
	}
	return &templateResponse, nil
}

func getAccessToken(openID string) (*GetAccessTokenResponse, error) {
	var accessTokenResponse GetAccessTokenResponse
	// 先从redis中拿
	accessToken, err := getAccessTokenFromRedis(openID)
	if accessToken != "" && err == nil {
		accessTokenResponse = GetAccessTokenResponse{AccessToken: accessToken}
		log.Info("从redis中获取到access_token", "access_token", accessToken)
		return &accessTokenResponse, nil
	}
	appID := config.ServerConfig.Appid
	secret := config.ServerConfig.Secret
	url := fmt.Sprintf(get_access_token_url, appID, secret)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("获取access_toke网络异常", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		log.Error("解析AccessToken失败", "err", err.Error())
		return nil, err
	}
	// 存到redis中
	if _, err := setAccessTokenToRedis(openID, accessTokenResponse.AccessToken); err != nil {
		log.Error("将access_token存储到redis中失败", "err", err.Error())
	}
	return &accessTokenResponse, nil
}

// 从redis中取access_token
func getAccessTokenFromRedis(openID string) (string, error) {
	key := fmt.Sprintf(constant.Redis_Prefix_Access_Token, openID)
	accessToken, err := db.Redis.Get(key).Result()
	return accessToken, err
}

// 将access_token存储到redis中
func setAccessTokenToRedis(openID, token string) (string, error) {
	key := fmt.Sprintf(constant.Redis_Prefix_Access_Token, openID)
	b, err := db.Redis.Set(key, token, 7200*time.Second).Result()
	return b, err
}

type TemplateMsg struct {
	Touser      string        `json:"touser"`      //接收者的OpenID
	TemplateID  string        `json:"template_id"` //模板消息ID
	FormID      string        `json:"form_id"`
	URL         string        `json:"url"`         //点击后跳转链接
	Miniprogram Miniprogram   `json:"miniprogram"` //点击跳转小程序
	Data        *TemplateData `json:"data"`
}
type Miniprogram struct {
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type TemplateData struct {
	First    KeyWordData `json:"first,omitempty"`
	Keyword1 KeyWordData `json:"keyword1,omitempty"`
	Keyword2 KeyWordData `json:"keyword2,omitempty"`
	Keyword3 KeyWordData `json:"keyword3,omitempty"`
	Keyword4 KeyWordData `json:"keyword4,omitempty"`
	Keyword5 KeyWordData `json:"keyword5,omitempty"`
}

type KeyWordData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type SendTemplateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MsgID   string `json:"msgid"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
