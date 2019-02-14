package controller

import (
	"fmt"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/constant"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/send", sendTemplate)
}

func sendTemplate(r *ghttp.Request) {
	templateMsg := &utils.TemplateMsg{}
	tempData := &utils.TemplateData{}
	tempData.First.Value = "测试模板消息"
	tempData.Keyword1.Value = "大家记得买票啊"
	tempData.Keyword2.Value = "马上就要放假了，大家记得买回家的票啊"
	tempData.Keyword3.Value = "2018-12-30 15:59"
	tempData.Keyword4.Value = "派大星"
	tempData.Keyword5.Value = "记得按时完成"
	templateMsg.Data = tempData
	formID := r.GetString("formID")
	log.Info("formID", "formID", formID)
	templateMsg.FormID = formID
	openID, _ := middleware.GetOpenID(r)
	templateMsg.Touser = openID
	templateMsg.TemplateID = constant.Tmeplate_Receive_Task_ID
	response, err := utils.SendTemplate(templateMsg)
	if err != nil {
		fmt.Println("发送模板消息失败", err.Error())
	}
	r.Response.WriteJson(response)
}
