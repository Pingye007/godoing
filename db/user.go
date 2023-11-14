package db

import (
	. "xorm.io/builder"
)

type User struct {
	UserId   int `xorm:"pk"`
	Password string
	UserName string
	Email    string
	Role     int
}

func (user *User) TableName() string {
	return tableName
}

const (
	tableName = "gd_user"
)

func QueryUserById(id int) (*User, error) {
	var user User
	_, err := Engine.ID(id).Get(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
