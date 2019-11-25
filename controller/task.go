package controller

import (
	"encoding/json"
	"time"

	"github.com/pibigstar/go-todo/constant"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/task/create", createTask)
	s.BindHandler("/task/list", listTasks)
	s.BindHandler("/task/changeStatus", changeStatus)
	s.BindHandler("/task/get", getTask)
	s.BindHandler("/task/getTaskData", getTaskData)
}

type Exerciser struct {
	userOpenID string `json:"userOpenId"`
}

type AppointTo struct {
	IsAll     bool   `json:"isAll"`
	Exerciser string `json:"exerciser"`
}

type CreateTaskRequest struct {
	TaskTitle      string    `json:"taskTitle"`
	TaskContent    string    `json:"taskContent"`
	TaskHTML       string    `json:"TaskHTML"`
	Assign         string    `json:"assign"`
	CompletionTime time.Time `json:"completionTime"`
	GroupID        int       `json:"groupId"`
	GroupName      string    `json:"groupName"`
	Tips           string    `json:"tips"`
	IsRemind       bool      `json:"isRemind"`
	IsAll          bool      `json:"isAll"`
	RemindAfterFin bool      `json:"remindAfterFin"`
	FileIds        []string  `json:"fileIds"`
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

type GetTaskRequest struct {
	Id int `json:"id"`
}

type GetTaskResponse struct {
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	HTML           string   `json:"html"`
	GroupName      string   `json:"groupName"`
	UserName       string   `json:"userName"`
	CreateTime     string   `json:"createTime"`
	CompletionTime string   `json:"completionTime"`
	FileIds        []string `json:"fileIds"`
}

type GetTaskDataResponse struct {
	Todo  int `json:"todo"`
	Doing int `json:"doing"`
	Done  int `json:"done"`
}

func createTask(r *ghttp.Request) {
	createTaskRequest := &CreateTaskRequest{}
	r.GetToStruct(createTaskRequest)

	mCreateTask := convertCreateTaskRequestToModel(createTaskRequest)
	openID, _ := middleware.GetOpenID(r)
	mCreateTask.CreateUser = openID
	if createTaskRequest.IsAll {
		groupUsers, err := models.MGroupUser.GetUsers(mCreateTask.GroupID)
		if err != nil {
			log.Error("get users by group id is failed")
		}
		for _, user := range groupUsers {
			mCreateTask.AppointTo = user.UserID
			err = models.MTask.Create(mCreateTask)
			if err != nil {
				log.Error("create task is failed", "user openId", user.UserID)
			}
		}
	} else {
		mCreateTask.AppointTo = createTaskRequest.Assign
		err := models.MTask.Create(mCreateTask)
		if err != nil {
			log.Error("create task is failed")
			r.Response.WriteJson(utils.ErrorResponse("创建任务失败"))
			r.Exit()
		}
	}
	// 开启定时提醒
	if mCreateTask.IsRemind {
		go sendTemplateMsg(mCreateTask, createTaskRequest.IsAll)
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

func getTask(r *ghttp.Request) {
	getTaskRequest := new(GetTaskRequest)
	r.GetToStruct(getTaskRequest)

	task, err := models.MTask.GetTask(getTaskRequest.Id)
	if err != nil {
		log.Error("get task is failed", "taskId", getTaskRequest.Id)
	}
	response := convertTaskToResponse(task)
	user, err := models.MUser.GetUserByOpenID(task.CreateUser)
	if err != nil {
		log.Error("get user is failed", "userOPenId", task.CreateUser)
	} else {
		response.UserName = user.NickName
	}
	// set isRead is true
	err = models.MTask.SetRead(getTaskRequest.Id)
	if err != nil {
		log.Error("set read is failed", "taskId", getTaskRequest.Id)
	}
	r.Response.WriteJson(utils.SuccessWithData("OK", response))
}

func getTaskData(r *ghttp.Request) {
	openId, err := middleware.GetOpenID(r)
	if err != nil {
		r.Response.WriteJson("请先登录")
		r.Exit()
	}
	todo, err := models.MTask.CountTask(openId, constant.TODO)
	doing, err := models.MTask.CountTask(openId, constant.DOING)
	done, err := models.MTask.CountTask(openId, constant.DONE)
	response := &GetTaskDataResponse{
		Todo:  todo,
		Doing: doing,
		Done:  done,
	}
	r.Response.WriteJson(utils.SuccessWithData("OK", response))
}

func convertTaskToResponse(task *models.Task) *GetTaskResponse {
	var fileIds []string
	if task.FileIds != "" {
		err := json.Unmarshal([]byte(task.FileIds), &fileIds)
		if err != nil {
			log.Error("Failed to unmarshal file ids", "err", err.Error())
		}
	}
	return &GetTaskResponse{
		Title:          task.TaskTitle,
		Content:        task.TaskContent,
		GroupName:      task.GroupName,
		HTML:           task.TaskHTML,
		CompletionTime: utils.TimeFormat(task.CompletionTime),
		CreateTime:     utils.TimeFormat(task.CreateTime),
		FileIds:        fileIds,
	}
}

func sendTemplateMsg(task *models.Task, isAll bool) {
	user, _ := models.MUser.GetUserByOpenID(task.CreateUser)
	userName := user.RealName
	if userName == "" {
		userName = user.NickName
	}
	templateMsg := &utils.TemplateMsg{}
	tempData := &utils.TemplateData{}
	tempData.Keyword1.Value = task.TaskTitle
	tempData.Keyword2.Value = task.TaskContent
	tempData.Keyword3.Value = utils.TimeFormat(task.CompletionTime)
	tempData.Keyword4.Value = userName
	tempData.Keyword5.Value = task.Tips
	templateMsg.Data = tempData
	templateMsg.Touser = user.OpenID
	templateMsg.TemplateID = constant.TemplateReceiveTaskId
	// 获取formID
	// 所有人
	if isAll {
		openIds, err := models.MGroupUser.GetUserOpenIDs(task.GroupID)
		if err != nil {
			log.Error("get user openIds is failed", "err", err.Error())
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
		openIds = append(openIds, task.AppointTo)
		formIds := models.MGroupUser.GetFormIds(openIds)
		for _, formId := range formIds {
			templateMsg.FormID = formId
			go utils.SendTemplate(templateMsg)
		}
	}
}

func convertCreateTaskRequestToModel(request *CreateTaskRequest) *models.Task {
	var fileIds string
	if len(request.FileIds) > 0 {
		bytes, err := json.Marshal(request.FileIds)
		if err != nil {
			log.Error("Failed to marshal file ids", "err", err.Error())
		} else {
			fileIds = string(bytes)
		}
	}
	return &models.Task{
		TaskTitle:      request.TaskTitle,
		TaskContent:    request.TaskContent,
		TaskHTML:       request.TaskHTML,
		AppointTo:      request.Assign,
		CompletionTime: request.CompletionTime,
		IsRemind:       request.IsRemind,
		Tips:           request.Tips,
		GroupID:        request.GroupID,
		GroupName:      request.GroupName,
		IsDelete:       false,
		IsRead:         false,
		CreateTime:     time.Now(),
		FileIds:        fileIds,
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
