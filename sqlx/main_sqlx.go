package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sdm_demo_todolist/shared"
	"sdm_demo_todolist/sqlx/dbal"
	"sdm_demo_todolist/sqlx/handlers"
	"sdm_demo_todolist/sqlx/swagger"
)

// @title			Todolist API
// @version		0.0.1
// @description	Todolist API
// @in				header
// @BasePath		/api
// @accept			json
// @produce		json
// @host			127.0.0.1:8080
// @schemes		http
func main() {
	err := dbal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		_ = dbal.CloseDB()
	}()

	gin.SetMode(gin.ReleaseMode)

	myRouter := gin.New()

	swagger.Init(myRouter)

	//myOS, myArch := runtime.GOOS, runtime.GOARCH
	//inContainer := " docker,"
	//if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
	//	inContainer = ""
	//}

	whoIam := fmt.Sprintf(`sqlx, sqlite3`)

	shared.AssignHandlers(myRouter, whoIam, handlers.NewProjectHandlers(), handlers.NewTaskHandlers())

	log.Fatal(myRouter.Run(":8080"))
}
