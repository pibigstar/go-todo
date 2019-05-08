package admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/api/task/list", taskList)
}

func taskList(r *ghttp.Request) {
	tasks, err := models.MTask.TaskList()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, tasks)
}


