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
}

func groupList(r *ghttp.Request) {
	groups, err := models.MGroup.ListGroup()
	if err != nil {
		utils.Error(r)
	}
	utils.Success(r, groups)
}


