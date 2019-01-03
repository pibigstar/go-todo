package controller

import (
	"encoding/json"
	"time"

	"github.com/pibigstar/go-todo/constant"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/task/create", createTask)
}

type Exerciser struct {
	userOpenID string `json:"userOpenId"`
}

type AppointTo struct {
	IsAll bool `json:"isAll"`
	Exercisers []Exerciser `json:"exercisers"`
}

type CreateTaskRequest struct {
	TaskTitle      string    `json:"taskTitle"`
	TaskContent    string    `json:"taskContent"`
	AppointTo      string    `json:"appointTo"`
	CompletionTime time.Time `json:"completionTime"`
	GroupID        int       `json:"group_id"`
	Tips           string    `json:"tips"`
	IsRemind       bool      `json:"isRemind"`
}

func createTask(r *ghttp.Request) {
	createTaskRequest := new(CreateTaskRequest)
	r.GetJson().ToStruct(createTaskRequest)
	middleware.CheckToken(r)
	mCreateTask := convertCreateTaskRequestToModel(createTaskRequest)
	openID, _ := middleware.GetOpenID(r)
	mCreateTask.CreateUser = openID
	taskId := models.MTask.Create(mCreateTask)
	if taskId == 0 {
		log.Error("创建任务失败")
		r.Response.WriteJson(utils.ErrorResponse("创建任务失败"))
		r.Exit()
	}
	if mCreateTask.IsRemind {
		go sendTemplateMsg(mCreateTask)
	}
	r.Response.WriteJson(utils.SuccessResponse("创建成功"))
}

func sendTemplateMsg(task *models.Task) {
	user, _ := models.MUser.GetUserByOpenID(task.CreateUser)
	userName := user.RealName
	if userName == "" {
		userName = user.NickName
	}
	templateMsg := &utils.TemplateMsg{}
	tempData := &utils.TemplateData{}
	tempData.Keyword1.Value = task.TaskTitle
	tempData.Keyword2.Value = task.TaskContent
	tempData.Keyword3.Value = task.CompletionTime.Format("2006-01-02 15:04:05")
	tempData.Keyword4.Value = userName
	tempData.Keyword5.Value = task.Tips
	templateMsg.Data = tempData
	templateMsg.Touser = user.OpenID
	templateMsg.TemplateID = constant.Tmeplate_Receive_Task_ID
	// 获取formID
	data := []byte(task.AppointTo)
	var appointTo AppointTo
	err := json.Unmarshal(data, &appointTo)
	if err!=nil {
		log.Error("解析指派人出错","err",err.Error())
	}
	// 所有人
	if appointTo.IsAll {
		openIds, err := models.MGroupUser.GetUserOpenIDs(task.GroupID)
		if err != nil {
			log.Error("获取群成员OpenID失败","err",err.Error())
		}
		if len(openIds) > 0 {
			formIds := models.MGroupUser.GetFormIds(openIds)
			for _,formId := range formIds {
				templateMsg.FormID = formId
				go utils.SendTemplate(templateMsg)
			}
		}
	} else {
		// 指定人
		var openIds []string
		for _,exerciser := range appointTo.Exercisers {
			openIds = append(openIds, exerciser.userOpenID)
		}
		formIds := models.MGroupUser.GetFormIds(openIds)
		for _,formId := range formIds {
			templateMsg.FormID = formId
			go utils.SendTemplate(templateMsg)
		}
	}
}

func convertCreateTaskRequestToModel(request *CreateTaskRequest) *models.Task {
	return &models.Task{
		TaskTitle:      request.TaskTitle,
		TaskContent:    request.TaskContent,
		AppointTo:      request.AppointTo,
		CompletionTime: request.CompletionTime,
		IsRemind:       request.IsRemind,
		Tips:           request.Tips,
		GroupID:        request.GroupID,
		IsDelete:       false,
		CreateTime:     time.Now(),
	}
}
