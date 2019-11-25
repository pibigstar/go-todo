package test

import (
	"testing"

	"github.com/pibigstar/go-todo/utils"
)

func TestUUID(t *testing.T) {
	uid := utils.GetUUID()
	t.Log(uid)
}
