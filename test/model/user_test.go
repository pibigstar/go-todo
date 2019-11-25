package test

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/test/gtest"
	"testing"

	"github.com/pibigstar/go-todo/models"
)

func TestCreateUser(t *testing.T) {
	user := &models.User{
		OpenID:   "pibigstar",
		NickName: "派大星",
	}
	err := user.Create(user)
	gtest.Assert(err, nil)

	getUser, err := models.MUser.GetUserByOpenID("pibigstar")
	gtest.Assert(err, nil)
	glog.Print(getUser)

}
