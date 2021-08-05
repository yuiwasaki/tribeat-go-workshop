package main

import (
	"fmt"

	"github.com/yuiwasaki/tribeat-go-workshop/clog"
	"github.com/yuiwasaki/tribeat-go-workshop/models"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

func transaction() (err error) {
	fmt.Println("START")
	model, err := models.NewModel()
	if err != nil {
		clog.Println(err)
		return err
	}
	tx := model.Begin()
	defer func() {
		if err != nil {
			model.Rollback()
			clog.Println("ROLEBACK")
		}
	}()
	rtn := tx.Create(models.Group{
		Group: oapi.Group{
			Id: "hello",
		},
		DefaultModel: models.InitDefaultModel(),
	})
	if rtn.Error != nil {
		clog.Println(rtn.Error.Error())
		return rtn.Error
	}
	rtn = tx.Create(models.Group{
		Group: oapi.Group{
			Id: "hello",
		},
		DefaultModel: models.InitDefaultModel(),
	})
	if rtn.Error != nil {
		clog.Println(rtn.Error.Error())
		return rtn.Error
	}
	model.Commit()
	clog.Println("END")
	return nil //fmt.Errorf("hello")
}

func main() {
	transaction()
}
