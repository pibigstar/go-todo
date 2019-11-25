package test

import (
	"github.com/pibigstar/go-todo/utils"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	t.Log(utils.TimeFormat(time.Now()))
}
