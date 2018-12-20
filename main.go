package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pibigstar/go-todo/router"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("main")

func main() {

	app := gin.New()

	router.Route(app)

	app.Run(":8023")
}
