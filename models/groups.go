package models

import (
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
	tx := model.Create(group)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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
