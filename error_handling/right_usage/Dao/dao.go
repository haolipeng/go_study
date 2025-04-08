package Dao

import (
	"fmt"
)

var ErrDBConnection = fmt.Errorf("database connection failed")

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) GetUserByID(id int) (string, error) {
	// 模拟数据库错误
	if id < 0 {
		// 使用fmt.Errorf和%w包装错误
		return "", fmt.Errorf("dao: database lookup failed: %w", ErrDBConnection)
	}
	return "User Name", nil
}
