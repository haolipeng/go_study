package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //测试：断言 DB.Get()方法是否被调用

	m := NewMockDB(ctrl) //定义在db_mock.go中，由mockgen自动生成
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1,but got", v)
	}
}

func TestGetAnyFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //测试：断言 DB.Get()方法是否被调用

	m := NewMockDB(ctrl) //定义在db_mock.go中，由mockgen自动生成
	m.EXPECT().Get(gomock.Any()).Return(630, nil)

	if v := GetFromDB(m, "haolipeng"); v != -1 {
		t.Error("expected -1,but got", v)
		t.Fatal("expected -1,but got", v)
	}
}
