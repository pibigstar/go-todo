package test

import (
	"fmt"
	"testing"

	"github.com/pibigstar/go-todo/utils"
)

func TestMd5(t *testing.T) {
	fmt.Println(utils.Md5("pibigstar"))
}
