package controller

import (
	"time"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/group/create", createGroup)
	s.BindHandler("/group/:userID/list", getGroupsByUserID)
	s.BindHandler("/group/join", joinGroup)
}

// CreateGroupRequest 创建组织请求体
type CreateGroupRequest struct {
	GroupName     string `json:"group_name"`
	GroupDescribe string `json:"group_describe"`
	GroupMaster   string `json:"group_master"`
	GroupCode     string `json:"group_code"`
}

// JoinGroupRequest 加入组织请求体
type JoinGroupRequest struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

// createGroup 创建组织
func createGroup(r *ghttp.Request) {
	createGroupRequest := new(CreateGroupRequest)
	r.GetPostToStruct(createGroupRequest)

	err := models.MGroup.Create(converCreateGroupToModel(createGroupRequest))
	if err != nil {
		r.Response.WriteJson(errorResponse(err.Error()))
	}
	r.Response.WriteJson(successResponse("ok"))

}

// getGroupsByUserID 获取用户所有组织
func getGroupsByUserID(r *ghttp.Request) {
	userID := r.GetInt("userID")
	groups, err := models.MGroup.GetGroupByID(userID)
	if err != nil {
		r.Response.WriteJson(errorResponse(err.Error()))
	}
	r.Response.WriteJson(successWithData("ok", groups))
}

// joinGroup 加入组织
func joinGroup(r *ghttp.Request) {

	joinGroupRequest := new(JoinGroupRequest)
	r.GetPostToStruct(joinGroupRequest)

	groupUser := converJoinGroupToModel(joinGroupRequest)

	err := models.MGroupUser.Create(groupUser)
	if err != nil {
		r.Response.WriteJson(errorResponse(err.Error()))
	}
	r.Response.WriteJson(successResponse("ok"))
}

func converCreateGroupToModel(createGroup *CreateGroupRequest) *models.Group {
	groupCode := utils.Md5(createGroup.GroupCode)
	return &models.Group{
		GroupName:     createGroup.GroupName,
		GroupDescribe: createGroup.GroupDescribe,
		GroupMaster:   createGroup.GroupMaster,
		GroupCode:     string(groupCode),
		IsDelete:      false,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
}

func converJoinGroupToModel(request *JoinGroupRequest) *models.GroupUser {
	return &models.GroupUser{
		GroupID:    request.GroupID,
		UserID:     request.UserID,
		CreateTime: time.Now(),
		IsDelete:   false,
	}
}
