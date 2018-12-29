package controller

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/signature", signature)
}

var (
	token          = "pibigstar123"
	EncodingAESKey = "BguNrFfYAciHTeeR0NueU6yqIeOrpix8oc7Pl1sPZVH"
)

type GetSignature struct {
	Signature string `json:"signature"` //微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
	Timestamp string `json:"timestamp"` //时间戳
	Nonce     string `json:"nonce"`     //随机数
	Echostr   string `json:"echostr"`   //随机字符串
}

func signature(r *ghttp.Request) {
	getSignature := new(GetSignature)
	r.GetToStruct(getSignature)
	sign := getSignature.Signature
	nonce := getSignature.Nonce
	timestamp := getSignature.Timestamp
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)

	tempStr := fmt.Sprintf("%s%s%s", strs[0], strs[1], strs[2])
	fmt.Println(tempStr)
	fmt.Println("sign", sign)
	h := sha1.New()
	io.WriteString(h, tempStr)
	result := fmt.Sprintf("%x\n", h.Sum(nil))
	if result != sign {
		r.Response.Write(false)
	}
	fmt.Println("result", result)
	r.Response.Write(true)
}
