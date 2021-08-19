package main

import (
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/d210819/router"
)

func main() {
	e := echo.New()
	f, err := middleware.OapiValidatorFromYamlFile("./sample.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	e.Use(f)
	handler := router.NewAPIHandler()
	handler.RegisterHandler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 3003)))
}
