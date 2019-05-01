package test

import (
	"fmt"
	"github.com/pibigstar/go-todo/utils"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	fmt.Print(utils.TimeFormat(time.Now()))
}
