package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenderCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
