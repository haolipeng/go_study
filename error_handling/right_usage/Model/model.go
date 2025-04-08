package Model

import (
	"fmt"
	"go_study/error_handling/right_usage/Dao"
)

// UserModel Model layer
type UserModel struct {
	dao *Dao.UserDAO
}

func NewUserModel(dao *Dao.UserDAO) *UserModel {
	return &UserModel{dao: dao}
}

func (m *UserModel) GetUserProfile(id int) (string, error) {
	username, err := m.dao.GetUserByID(id)
	if err != nil {
		// 包装错误，添加上下文信息
		return "", fmt.Errorf("model: failed to get user profile for ID %d: %w", id, err)
	}
	return username, nil
}
