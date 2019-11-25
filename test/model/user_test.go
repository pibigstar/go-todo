package test

import (
	"context"
	"testing"

	"github.com/pibigstar/go-todo/models"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	user := &models.User{
		OpenID:   "pibigstar",
		NickName: "派大星",
	}
	user.Create(user)

	getUser, err := models.MUser.GetUserByOpenID("pibigstar")
	if err != nil {
		log.CtxError(ctx, "获取用户信息失败", "err", err.Error())
	}

	log.CtxInfo(ctx, "获取用户信息成功", "user", getUser)

}
