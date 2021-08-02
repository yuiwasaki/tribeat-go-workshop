package models

import (
	"testing"
)

func TestGroups(t *testing.T) {
	model, err := NewModel()
	if err != nil {
		t.Error(err)
		return
	}
	gModel := NewGroupModel(model)
	groups, err := gModel.All()
	if err != nil {
		t.Error(err)
		return
	}
	if len(groups) != 0 {
		t.Error(groups)
		return
	}
	group := Group{}
	group.Id = "HOGE"
	group.Name = "PIYO"
	err = gModel.Save(group)
	if err != nil {
		t.Error(err)
	}
	defer gModel.Delete("HOGE")
	if err != nil {
		t.Error(err)
	}
}
