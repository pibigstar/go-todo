package controller

import (
	"encoding/json"
	"time"

	"github.com/pibigstar/go-todo/constant"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/task/create", createTask)
	s.BindHandler("/task/list", listTasks)
	s.BindHandler("/task/changeStatus", changeStatus)
}

type Exerciser struct {
	userOpenID string `json:"userOpenId"`
}

type AppointTo struct {
	IsAll      bool        `json:"isAll"`
	Exercisers []Exerciser `json:"exercisers"`
}

type CreateTaskRequest struct {
	TaskTitle      string    `json:"taskTitle"`
	TaskContent    string    `json:"taskContent"`
	AppointTo      string    `json:"appointTo"`
	CompletionTime time.Time `json:"completionTime"`
	GroupID        int       `json:"groupId"`
	GroupName      string    `json:"groupName"`
	Tips           string    `json:"tips"`
	IsRemind       bool      `json:"isRemind"`
	RemindAfterFin bool      `json:"remindAfterFin"`
}

type ListTaskRequest struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

type ListTaskResponse struct {
	Tasks     []Tasks `json:"tasks"`
	UnReadNum int     `json:"unReadNum"`
}

type Tasks struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Time      string `json:"time"`
	IsRead    bool   `json:"isRead"`
	GroupName string `json:"groupName"`
}

type ChangeStatusRequest struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

func createTask(r *ghttp.Request) {
	createTaskRequest := new(CreateTaskRequest)
	r.GetJson().ToStruct(createTaskRequest)
	mCreateTask := convertCreateTaskRequestToModel(createTaskRequest)
	openID, _ := middleware.GetOpenID(r)
	mCreateTask.CreateUser = openID
	err := models.MTask.Create(mCreateTask)
	if err != nil {
		log.Error("创建任务失败")
		r.Response.WriteJson(utils.ErrorResponse("创建任务失败"))
		r.Exit()
	}
	// 开启定时提醒
	if mCreateTask.IsRemind {
		go sendTemplateMsg(mCreateTask)
	}
	r.Response.WriteJson(utils.SuccessResponse("创建成功"))
}

func listTasks(r *ghttp.Request) {
	listTaskRequest := new(ListTaskRequest)
	r.GetToStruct(listTaskRequest)

	openId, err := middleware.GetOpenID(r)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse("user is not login"))
		r.Exit()
	}
	tasks, err := models.MTask.ListTask(openId, listTaskRequest.Status, listTaskRequest.Title)
	if err != nil {
		log.Error("list tasks is failed", "openId", openId, "status", listTaskRequest.Status)
	}
	response := convertListTaskToResponse(tasks)
	r.Response.WriteJson(utils.SuccessWithData("OK", response))
}

func changeStatus(r *ghttp.Request) {
	changeStatusRequest := new(ChangeStatusRequest)
	r.GetToStruct(changeStatusRequest)

	err := models.MTask.ChangeStatus(changeStatusRequest.Id, changeStatusRequest.Status)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse("update status is failed"))
		r.Exit()
	}
	r.Response.WriteJson(utils.SuccessResponse("OK"))
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
	if err != nil {
		log.Error("解析指派人出错", "err", err.Error())
	}
	// 所有人
	if appointTo.IsAll {
		openIds, err := models.MGroupUser.GetUserOpenIDs(task.GroupID)
		if err != nil {
			log.Error("获取群成员OpenID失败", "err", err.Error())
		}
		if len(openIds) > 0 {
			formIds := models.MGroupUser.GetFormIds(openIds)
			for _, formId := range formIds {
				templateMsg.FormID = formId
				go utils.SendTemplate(templateMsg)
			}
		}
	} else {
		// 指定人
		var openIds []string
		for _, exerciser := range appointTo.Exercisers {
			openIds = append(openIds, exerciser.userOpenID)
		}
		formIds := models.MGroupUser.GetFormIds(openIds)
		for _, formId := range formIds {
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
		GroupName:      request.GroupName,
		IsDelete:       false,
		IsRead:         false,
		CreateTime:     time.Now(),
	}
}

func convertListTaskToResponse(tasks []models.Task) *ListTaskResponse {
	var taskResponse []Tasks
	var unReadNum = 0
	for _, task := range tasks {
		newTask := Tasks{
			Id:        task.ID,
			Title:     task.TaskTitle,
			Content:   task.TaskContent,
			Time:      utils.DateFormat(task.CompletionTime),
			IsRead:    task.IsRead,
			GroupName: task.GroupName,
		}
		if !task.IsRead {
			unReadNum++
		}
		taskResponse = append(taskResponse, newTask)
	}
	return &ListTaskResponse{
		Tasks:     taskResponse,
		UnReadNum: unReadNum,
	}
}
