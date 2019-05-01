package controller

import (
	"time"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/group/create", createGroup)
	s.BindHandler("/group/list", listGroups)
	s.BindHandler("/group/join", joinGroup)
	s.BindHandler("/group/search", searchGroup)
	s.BindHandler("/group/members", getMembers)
}

// createGroup 创建组织
func createGroup(r *ghttp.Request) {
	createGroupRequest := new(CreateGroupRequest)
	r.GetJson().ToStruct(createGroupRequest)
	mCreateGroup := convertCreateGroupToModel(createGroupRequest)
	openID, err := middleware.GetOpenID(r)
	mCreateGroup.GroupMaster = openID
	err = models.MGroup.Create(mCreateGroup)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
	}
	r.Response.WriteJson(utils.SuccessResponse("ok"))

}

// listGroups 获取用户创建的群
func listGroups(r *ghttp.Request) {
	userID, _ := middleware.GetOpenID(r)
	createGroups, err := models.MGroup.GetUserCreateGroups(userID)
	if err != nil {
		log.Error("获取用户创建的群失败", "err", err.Error())
	}
	joinGroups, err := models.MGroupUser.GetUserJoinGroups(userID)
	if err != nil {
		log.Error("获取用户加入的群失败", "err", err.Error())
	}
	getGroupsResponse := &GetGroupResponse{
		CreateGroups: createGroups,
		JoinGroups:   joinGroups,
	}
	r.Response.WriteJson(utils.SuccessWithData("ok", getGroupsResponse))
}

// searchGroup 查询组织
func searchGroup(r *ghttp.Request) {
	groupID := r.GetInt("groupId")
	group, err := models.MGroup.GetGroupByID(groupID)
	if err != nil {
		log.Error(err.Error(), "groupID", groupID)
		r.Response.WriteJson(utils.ErrorResponse("查询组织失败"))
		r.Exit()
	}
	groupResponse := convertModelGroupToResponse(group)
	r.Response.WriteJson(utils.SuccessWithData("ok!", groupResponse))
}

// joinGroup 加入组织
func joinGroup(r *ghttp.Request) {
	joinGroupRequest := new(JoinGroupRequest)
	r.GetJson().ToStruct(joinGroupRequest)
	group, err := models.MGroup.GetGroupByID(joinGroupRequest.GroupID)
	if err != nil {
		log.Error("没有此组织", "groupID", joinGroupRequest.GroupID)
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
	switch group.JoinMethod {
	// 秘钥
	case "1":
		if joinGroupRequest.GroupCode != group.GroupCode {
			r.Response.WriteJson(utils.ErrorResponse("秘钥错误，请联系组织创建者"))
			r.Exit()
		}
	// 回答问题
	case "2":
		if joinGroupRequest.Answer != group.Answer {
			r.Response.WriteJson(utils.ErrorResponse("答案错误，请联系组织创建者"))
			r.Exit()
		}
	}
	groupUser := convertJoinGroupToModel(joinGroupRequest)
	groupUser.UserID, _ = middleware.GetOpenID(r)
	// 判断是否已经加入该组织了
	isExist, err := models.MGroupUser.IsExist(groupUser.UserID, group.ID)
	if isExist {
		r.Response.WriteJson(utils.ErrorResponse("你已加入该组织，请不要重复加入"))
		r.Exit()
	}
	// 获取用户名
	user, err := models.MUser.GetUserByOpenID(groupUser.UserID)
	if err != nil {
		log.Info("此用户名为空", "OpenId", groupUser.UserID)
	}
	groupUser.UserName = user.NickName
	err = models.MGroupUser.Create(groupUser)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
	}
	r.Response.WriteJson(utils.SuccessResponse("ok"))
}

// getMembers 获取此群下面的群成员
func getMembers(r *ghttp.Request) {
	getMemberRequest := new(GetMemberRequest)
	r.GetToStruct(getMemberRequest)
	groupUsers, err := models.MGroupUser.GetUsers(getMemberRequest.GroupId)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
	r.Response.WriteJson(utils.SuccessWithData("ok", groupUsers))
}

func convertCreateGroupToModel(createGroup *CreateGroupRequest) *models.Group {
	groupCode := utils.GetUUID()
	return &models.Group{
		GroupName:     createGroup.GroupName,
		GroupDescribe: createGroup.GroupDescribe,
		JoinMethod:    createGroup.JoinMethod,
		Question:      createGroup.Question,
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
		CreateTime: time.Now(),
		IsDelete:   false,
	}
}

func convertModelGroupToResponse(group *models.Group) *SearchGroupResponse {
	return &SearchGroupResponse{
		GroupName:     group.GroupName,
		GroupDescribe: group.GroupDescribe,
		GroupID:       group.ID,
		JoinMethod:    group.JoinMethod,
		Question:      group.Question,
	}
}

// CreateGroupRequest 创建组织请求体
type CreateGroupRequest struct {
	GroupName     string `json:"groupName"`
	GroupDescribe string `json:"groupDescribe"`
	GroupMaster   string `json:"groupMaster"`
	GroupCode     string `json:"groupCode"`
	JoinMethod    string `json:"joinMethod"`
	Question      string `json:"question"`
	Answer        string `json:"answer"`
}

// JoinGroupRequest 加入组织请求体
type JoinGroupRequest struct {
	GroupID    int    `json:"groupId"`
	GroupCode  string `json:"groupCode"`
	JoinMethod string `json:"joinMethod"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}

// GetGroupResponse 获取用户所有群响应体
type GetGroupResponse struct {
	CreateGroups *[]models.Group `json:"createGroups"`
	JoinGroups   *[]models.Group `json:"joinGroups"`
}

// SearchGroupResponse 搜索组织响应体
type SearchGroupResponse struct {
	GroupName     string `json:"groupName"`
	GroupDescribe string `json:"groupDescribe"`
	GroupID       int    `json:"groupId"`
	JoinMethod    string `json:"joinMethod"`
	Question      string `json:"question"`
}

type GetMemberRequest struct {
	GroupId int `json:"groupId"`
}
