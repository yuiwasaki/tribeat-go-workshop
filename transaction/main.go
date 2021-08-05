package main

import (
	"fmt"

	"github.com/yuiwasaki/tribeat-go-workshop/models"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

func transaction() (err error) {
	fmt.Println("START")
	model, err := models.NewModel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	tx := model.Begin()
	defer func() {
		if err != nil {
			model.Rollback()
			fmt.Println("ROLEBACK")
		}
	}()
	rtn := tx.Create(models.Group{
		Group: oapi.Group{
			Id: "hello",
		},
		DefaultModel: models.InitDefaultModel(),
	})
	if rtn.Error != nil {
		return rtn.Error
	}
	rtn = tx.Create(models.Group{
		Group: oapi.Group{
			Id: "hello",
		},
		DefaultModel: models.InitDefaultModel(),
	})
	if rtn.Error != nil {
		return rtn.Error
	}
	model.Commit()
	fmt.Println("END")
	return nil //fmt.Errorf("hello")
}

func main() {
	transaction()
}
