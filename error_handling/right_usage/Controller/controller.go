package Controller

import (
	"fmt"
	"go_study/error_handling/right_usage/Model"
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
		// 包装错误，添加上下文信息
		return "", fmt.Errorf("controller: API error when getting user ID %d: %w", id, err)
	}
	return username, nil
}
