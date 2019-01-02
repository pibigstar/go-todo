package test

import (
	"fmt"
	"github.com/pibigstar/go-todo/utils"
	"testing"
)

func TestGenderID(t *testing.T) {
	// goroutine 和 chan
	for {
		fmt.Println(utils.Id.Next())
	}
	// mysql
	//g:= wuid.NewWUID("default", nil)
	//g.LoadH24FromMysql("127.0.0.1:3306", "root", "root", "test", "wuid")
	//// Generate
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d\n", g.Next())
	//}
	// redis
	//g := wuid.NewWUID("default", nil)
	//g.LoadH24FromRedis("127.0.0.1:6379", "", "wuid")
	//
	// Generate
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%#016x\n", g.Next())
	//}

}
