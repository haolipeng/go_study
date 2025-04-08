package main

import (
	"errors"
	"fmt"
	"go_study/error_handling/wrong_usage/Controller"
	"go_study/error_handling/wrong_usage/Dao"
	"go_study/error_handling/wrong_usage/Model"
)

func main() {
	fmt.Println("=== 演示1: 错误的错误处理方式（既记录日志又返回错误）但不使用结构体内logger ===")
	dao := Dao.NewUserDAO()
	model := Model.NewUserModel(dao)
	controller := Controller.NewUserController(model)
	_, err := controller.HandleGetUser(-1)

	if err != nil {
		if err == Dao.ErrDBConnection {
			fmt.Println("✓ 能直接识别出原始错误")
		} else {
			fmt.Println("✗ 无法直接识别出原始错误（因为错误已被包装）")
		}

		// 尝试使用errors.Is
		fmt.Println("\n尝试使用errors.Is检查原始错误:")
		if errors.Is(err, Dao.ErrDBConnection) {
			fmt.Println("✓ 使用errors.Is能识别出原始错误")
		} else {
			fmt.Println("✗ 即使使用errors.Is也无法识别出原始错误")
		}
	}
}
