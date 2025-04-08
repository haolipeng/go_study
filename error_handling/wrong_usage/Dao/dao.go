package Dao

import (
	"fmt"
)

var ErrDBConnection = fmt.Errorf("database connection failed")

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// GetUserByID 最底层的错误
func (dao *UserDAO) GetUserByID(id int) (string, error) {
	// 模拟数据库错误
	if id < 0 {
		// 错误示例：先记录日志，然后返回错误
		fmt.Printf("[DAO] Failed to get user with ID %d: %v\n", id, ErrDBConnection)
		return "", ErrDBConnection
	}
	return "User Name", nil
}
