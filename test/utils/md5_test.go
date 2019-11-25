package test

import (
	"testing"

	"github.com/pibigstar/go-todo/utils"
)

func TestMd5(t *testing.T) {
	t.Log(utils.Md5("pibigstar"))
}
