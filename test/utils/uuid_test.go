package test

import (
	"fmt"
	"testing"

	"github.com/pibigstar/go-todo/utils"
)

func TestUUID(t *testing.T) {
	uid := utils.GetUUID()
	fmt.Println(uid)
}
