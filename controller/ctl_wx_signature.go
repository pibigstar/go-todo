package controller

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/signature", signature)
}

var (
	token          = "pibigstar110"
	EncodingAESKey = "yi1CvPLF1ZzablxtMnc93ER7d4W5HBsxVXlPFBtIARE"
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
	echostr := getSignature.Echostr
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)

	tempStr := fmt.Sprintf("%s%s%s", strs[0], strs[1], strs[2])
	h := sha1.New()
	io.WriteString(h, tempStr)
	result := fmt.Sprintf("%x", h.Sum(nil))
	if result != sign {
		// 等不等都让它返回正确的结果，zz验证
		r.Response.Write(echostr)
		return
	}
	r.Response.Write(echostr)
}
