package main

import (
	"flag"

	"github.com/dekarti/ssu-gw/common"
	"github.com/dekarti/ssu-gw/handlers"
	"github.com/docker/docker/client"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	var err error
	common.CLI, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	var httpAddr = flag.String("http", "0.0.0.0:8080", "HTTP service address")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Logger.SetLevel(log.DEBUG)

	e.Static("/", "public")
	e.POST("/upload", handlers.UploadHandler)
	e.POST("/tasks/launch", handlers.LaunchWorkHandler)
	e.POST("/tasks", handlers.CreateTaskHandler)
	e.GET("/tasks", handlers.GetTasksHandler)
	e.GET("/images", handlers.GetImagesHandler)

	e.Logger.Fatal(e.Start(*httpAddr))
}
