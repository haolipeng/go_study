package Model

import (
	"fmt"
	"go_study/error_handling/wrong_usage/Dao"
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
		// 错误示例：又记录日志又返回错误
		fmt.Printf("[MODEL] Failed to get user profile for ID %d: %v\n", id, err)
		return "", fmt.Errorf("user profile retrieval error: %w", err)
	}
	return username, nil
}
