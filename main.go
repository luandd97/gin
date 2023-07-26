package main

import (
	"diluan/config"
	"diluan/middleware"
	"diluan/routes"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VERSION = "0.13"
)

func main() {
	fmt.Println(os.Getenv("MONGODB_URL"))
	config.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc // -> this is setting the global timezone

	app := gin.New()

	app.Use(middleware.CORSMiddleware())
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Success")
	})
	router := app.Group("api/v1")
	routes.InitRoutes(router)
	app.Run(":9999")
}
