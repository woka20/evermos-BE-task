package main

import (
	"evermos-be-task/task-store/config"
	"evermos-be-task/task-store/databases"
	"evermos-be-task/task-store/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/kataras/iris/v12"
)

func main() {
	config.ConfigInit()

	apps := gin.Default()
	newHandler := handler.NewOrderHandler()

	cmdString := command()
	fmt.Println("command " + cmdString)

	if cmdString == "migrate" {
		var dbSvc = databases.NewDatabaseRepo()
		_ = dbSvc.DBInit()
		_ = dbSvc.Migrate()
	} else {

		var dbSvc = databases.NewDatabaseRepo()
		_ = dbSvc.DBInit()
		apps.POST("/order", newHandler.GenerateOrder)

		fmt.Println("Server Ready!")
		apps.Run(":" + config.PORT)
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
