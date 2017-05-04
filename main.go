package main

import (
	"flag"

	"github.com/dekarti/ssu-gw/handlers"
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

	handlers.RunContainer(cli)

}
