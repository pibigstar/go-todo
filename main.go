package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/router"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("main")

func main() {

	port := config.ServerConfig.Port

	app := gin.New()

	router.Route(app)

	app.Run(":" + fmt.Sprintf("%d", port))
}
