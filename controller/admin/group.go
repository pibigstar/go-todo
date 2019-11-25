package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/api/group/list", groupList)
	s.BindHandler("/api/group/delete", groupDelete)
}

func groupList(r *ghttp.Request) {
	groups, err := models.MGroup.ListGroup()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, groups)
}

func groupDelete(r *ghttp.Request) {
	request := &IDRequest{}
	r.GetToStruct(request)
	if request.ID == 0 {
		return
	}
	err := models.MGroup.GroupDelete(request.ID)
	if err != nil {
		log.Error("delete group failed")
	}
	utils.SuccessResponse("OK")
}
