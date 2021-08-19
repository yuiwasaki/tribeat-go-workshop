package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/d210819/router"
)

func main() {
	e := echo.New()
	handler := router.NewAPIHandler()
	handler.RegisterHandler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 3003)))
}
