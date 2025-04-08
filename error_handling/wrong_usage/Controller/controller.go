package Controller

import (
	"fmt"
	"go_study/error_handling/wrong_usage/Model"
)

// Controller层
type UserController struct {
	model *Model.UserModel
}

func NewUserController(model *Model.UserModel) *UserController {
	return &UserController{
		model: model,
	}
}

func (c *UserController) HandleGetUser(id int) (string, error) {
	username, err := c.model.GetUserProfile(id)
	if err != nil {
		// 错误示例：又记录日志又返回错误
		fmt.Printf("[CONTROLLER] API error when getting user ID %d: %v\n", id, err)
		return "", fmt.Errorf("API error: %w", err)
	}
	return username, nil
}

func (c *UserController) HandleGetUserGood(id int) (string, error) {
	username, err := c.model.GetUserProfile(id)
	if err != nil {
		// 错误示例：又记录日志又返回错误
		return "", fmt.Errorf("API error: %w", err)
	}
	return username, nil
}
