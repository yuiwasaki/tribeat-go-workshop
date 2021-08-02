package main

import "github.com/yuiwasaki/tribeat-go-workshop/models"

func main() {
	model, _ := models.NewModel()
	model.AutoMigrate(&models.Group{})
}
