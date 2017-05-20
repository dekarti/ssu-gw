package main

import (
	"flag"

	// "github.com/dekarti/ssu-gw/handlers"
	"github.com/dekarti/ssu-gw/models"
	"github.com/docker/docker/client"
	//	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

func main() {
	//var httpAddr = flag.String("http", "0.0.0.0:8080", "HTTP service address")
	flag.Parse()

	//e := echo.New()
	//e.Use(middleware.Logger())

	//e.GET("/hello", handlers.HelloHandler)

	//e.Logger.Fatal(e.Start(*httpAddr))

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	task1 := &models.Task{
		Name:           "task 1",
		Description:    "Task 1",
		DefaultInput:   "",
		ExpectedOutput: "",
	}

	work1 := &models.Work{
		DockerClient: cli,
		Task:         task1,
		Path:         "/Users/aarutyunyan/Desktop/ssu-fl/tasks",
		Command:      []string{"python", "exp_nfa.py"},
	}
	work1.LaunchWork()

	//	handlers.RunContainer(cli)

}
