package admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
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
	request := new(IDRequest)
	r.GetJson().ToStruct(request)
	if request.ID == 0 {
		return
	}
	err := models.MGroup.GroupDelete(request.ID)
	if err != nil {
		log.Error("delete group failed")
	}
	utils.SuccessResponse("OK")
}
