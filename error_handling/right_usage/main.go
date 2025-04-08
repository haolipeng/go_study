package main

import (
	"errors"
	"fmt"
	"go_study/error_handling/right_usage/Controller"
	"go_study/error_handling/right_usage/Dao"
	"go_study/error_handling/right_usage/Model"
)

// 封装一个更好的错误显示函数，更清晰地展示错误链
func betterUnwrap(err error) {
	fmt.Println("\n=== 错误链分析 ===")
	fmt.Println("错误链从最外层到最内层：")

	depth := 1
	currentErr := err
	//visited := make(map[error]bool) // 使用error类型作为key

	for currentErr != nil {
		// 检查是否已经处理过相同的错误
		/*
			if visited[currentErr] {
				fmt.Printf("⚠️ 第%d层: <检测到重复错误，终止解析>\n", depth)
				break
			}
			visited[currentErr] = true
		*/

		// 打印当前层错误
		fmt.Printf("🔍 第%d层: %v\n", depth, currentErr)

		// 解包错误，使用标准库的errors.Unwrap
		nextErr := errors.Unwrap(currentErr)
		if errors.Is(nextErr, currentErr) {
			fmt.Printf("⚠️ 第%d层: <无法继续解包，终止解析>\n", depth+1)
			break
		}
		currentErr = nextErr
		depth++
	}
}

// 演示自定义错误类型的使用
type CustomError struct {
	Msg     string
	Code    int
	OrigErr error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[错误码:%d] %s: %v", e.Code, e.Msg, e.OrigErr)
}

func (e *CustomError) Unwrap() error {
	return e.OrigErr
}

// 展示错误链中的每一层
func showErrorLayers(err error) {
	fmt.Println("\n=== 错误分析报告 ===")

	// 1. 打印简单错误消息
	fmt.Println("\n📝 错误摘要:")
	fmt.Printf("   %v\n", err)

	// 2. 测试errors.Is功能
	fmt.Println("\n🔍 错误类型检查:")
	if errors.Is(err, Dao.ErrDBConnection) {
		fmt.Println("   ✓ 成功识别出原始的数据库连接错误")
	} else {
		fmt.Println("   ✗ 无法识别出原始的数据库连接错误")
	}

	// 3. 使用更好的解包方法
	betterUnwrap(err)
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
