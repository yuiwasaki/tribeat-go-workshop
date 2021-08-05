package models

import (
	"fmt"
	"time"

	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

type Group struct {
	oapi.Group
	DefaultModel
}

type GroupModel struct {
	*Model
}

func NewGroupModel(model *Model) *GroupModel {
	return &GroupModel{
		model,
	}
}

// All 全て取得
func (model *GroupModel) All() ([]Group, error) {
	var groups []Group
	tx := model.Find(&groups)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return groups, nil
}

// Save 保存
func (model *GroupModel) Save(group Group) error {
	now := time.Now()
	group.CreatedAt = now
	group.UpdatedAt = now
	fmt.Println(group)
	tx := model.Create(group)
	return tx.Error
}

// Delete 削除
func (model *GroupModel) Delete(id string) error {
	tx := model.DB.Delete(Group{
		Group: oapi.Group{
			Id: id,
		},
	})
	return tx.Error
}
