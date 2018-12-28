package controller

import (
	"time"

	"github.com/pibigstar/go-todo/middleware"

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
	GroupName     string `json:"groupName"`
	GroupDescribe string `json:"groupDescribe"`
	GroupMaster   string `json:"groupMaster"`
	GroupCode     string `json:"groupCode"`
	JoinMethod    string `json:"joinMethod"`
	Questiong     string `json:"question"`
	Answer        string `json:"answer"`
}

// JoinGroupRequest 加入组织请求体
type JoinGroupRequest struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

// createGroup 创建组织
func createGroup(r *ghttp.Request) {
	createGroupRequest := new(CreateGroupRequest)
	r.GetJson().ToStruct(createGroupRequest)
	// 判断token是否有效
	middleware.CheckToken(r)
	mCreateGroup := convertCreateGroupToModel(createGroupRequest)
	openID, err := middleware.GetOpenID(r)
	mCreateGroup.GroupMaster = openID
	err = models.MGroup.Create(mCreateGroup)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
	}
	r.Response.WriteJson(utils.SuccessResponse("ok"))

}

// getGroupsByUserID 获取用户所有组织
func getGroupsByUserID(r *ghttp.Request) {
	userID := r.GetInt("userID")
	groups, err := models.MGroup.GetGroupByID(userID)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
	}
	r.Response.WriteJson(utils.SuccessWithData("ok", groups))
}

// joinGroup 加入组织
func joinGroup(r *ghttp.Request) {

	joinGroupRequest := new(JoinGroupRequest)
	r.GetJson().ToStruct(joinGroupRequest)

	groupUser := convertJoinGroupToModel(joinGroupRequest)

	err := models.MGroupUser.Create(groupUser)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
	}
	r.Response.WriteJson(utils.SuccessResponse("ok"))
}

func convertCreateGroupToModel(createGroup *CreateGroupRequest) *models.Group {
	groupCode := utils.GetUUID()
	return &models.Group{
		GroupName:     createGroup.GroupName,
		GroupDescribe: createGroup.GroupDescribe,
		JoinMethod:    createGroup.JoinMethod,
		Question:      createGroup.Questiong,
		Answer:        createGroup.Answer,
		GroupCode:     groupCode,
		IsDelete:      false,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
}

func convertJoinGroupToModel(request *JoinGroupRequest) *models.GroupUser {
	return &models.GroupUser{
		GroupID:    request.GroupID,
		UserID:     request.UserID,
		CreateTime: time.Now(),
		IsDelete:   false,
	}
}
