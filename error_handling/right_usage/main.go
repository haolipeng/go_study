package main

import (
	"errors"
	"fmt"
	"go_study/error_handling/right_usage/Controller"
	"go_study/error_handling/right_usage/Dao"
	"go_study/error_handling/right_usage/Model"
)

// 展示错误链中的每一层
func showErrorLayers(err error) {
	// 1. 打印简单错误消息
	fmt.Println("\n📝 错误原文:")
	fmt.Printf("   %v\n", err)

	// 2. 测试errors.Is功能
	fmt.Println("\n🔍 错误类型检查:")
	if errors.Is(err, Dao.ErrDBConnection) {
		fmt.Println("   ✓ 成功识别出原始的数据库连接错误")
	} else {
		fmt.Println("   ✗ 无法识别出原始的数据库连接错误")
	}

	// 3. 使用更好的解包方法
	fmt.Println("\n=== 错误链分析 ===")
	fmt.Println("错误链从最外层到最内层：")

	for depth := 1; err != nil; depth++ {
		fmt.Printf("🔍 第%d层: %v\n", depth, err)
		err = errors.Unwrap(err)
	}
}

func main() {
	fmt.Println("=== 演示: 使用Go 1.13+标准库错误处理 ===")
	dao := Dao.NewUserDAO()
	model := Model.NewUserModel(dao)
	controller := Controller.NewUserController(model)

	// 模拟调用
	username, err := controller.HandleGetUser(-1)
	if err != nil {
		// 展示错误链的各个层次
		showErrorLayers(err)
	} else {
		fmt.Printf("User: %s\n", username)
	}
}
