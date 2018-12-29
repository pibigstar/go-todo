package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pibigstar/go-todo/config"
)

// 发送模板消息

var (
	send_template_url    = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
	get_access_token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

func SendTemplate(msg *TemplateMsg) (*SendTemplateResponse, error) {
	msg.Miniprogram.AppID = config.ServerConfig.Appid
	accessToken, err := getAccessToken()
	if err != nil {
		log.Error("获取accessToken失败")
		return nil, err
	}
	url := fmt.Sprintf(send_template_url, accessToken.AccessToken)

	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("解析模板消息错误", "err", err.Error())
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

func getAccessToken() (*GetAccessTokenResponse, error) {
	openID := config.ServerConfig.Appid
	secret := config.ServerConfig.Secret
	url := fmt.Sprintf(get_access_token_url, openID, secret)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("获取access_toke网络异常", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var accessToken GetAccessTokenResponse
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		log.Error("解析AccessToken失败", "err", err.Error())
		return nil, err
	}
	return &accessToken, nil
}

type TemplateMsg struct {
	Touser      string        `json:"touser"`
	TemplateID  string        `json:"template_id"`
	FormID      string        `json:"form_id"`
	URL         string        `json:"url"`
	Miniprogram Miniprogram   `json:"miniprogram"`
	Data        *TemplateData `json:"data"`
}
type Miniprogram struct {
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type TemplateData struct {
	First    KeyWordData `json:"first"`
	Keyword1 KeyWordData `json:"keyword1"`
	Keyword2 KeyWordData `json:"keyword2"`
	Keyword3 KeyWordData `json:"keyword3"`
	Keyword4 KeyWordData `json:"keyword4"`
	Keyword5 KeyWordData `json:"keyword5"`
}

type KeyWordData struct {
	Value string `json:"value"`
	Color string `json:"color"`
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
