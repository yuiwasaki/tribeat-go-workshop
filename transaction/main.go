package main

import (
	"fmt"

	"github.com/yuiwasaki/tribeat-go-workshop/clog"
	"github.com/yuiwasaki/tribeat-go-workshop/models"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

func transaction1() (err error) {
	fmt.Println("START")
	model, err := models.NewModel()
	if err != nil {
		clog.Println(err)
		return err
	}
	// トランズアクション開始
	tx := model.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
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
		tx.Rollback()
		return rtn.Error
	}
	fmt.Println("1件目成功")
	rtn = tx.Create(models.Group{
		Group: oapi.Group{
			Id: "hello",
		},
		DefaultModel: models.InitDefaultModel(),
	})
	if rtn.Error != nil {
		clog.Println(rtn.Error.Error())
		tx.Rollback()
		return rtn.Error
	}
	fmt.Println("2件目成功")
	tx.Commit()
	fmt.Println("END")
	return nil //fmt.Errorf("hello")
}

func main() {
	transaction1()
}
