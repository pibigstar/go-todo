package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/api/task/list", taskList)
	s.BindHandler("/api/task/delete", taskDelete)
}

type DeleteTaskRequest struct {
	ID int `json:"id"`
}

func taskList(r *ghttp.Request) {
	tasks, err := models.MTask.TaskList()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, tasks)
}

func taskDelete(r *ghttp.Request) {
	request := &DeleteTaskRequest{}
	r.GetToStruct(request)
	if request.ID == 0 {
		return
	}
	err := models.MTask.TaskDelete(request.ID)
	if err != nil {
		log.Error("delete task failed")
	}
	utils.SuccessResponse("OK")
}
