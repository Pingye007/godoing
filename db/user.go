package db

import (
	"github.com/Pingye007/godoing/log"
	. "xorm.io/builder"
)

type User struct {
	UserId   int
	Password string
	UserName string
	Email    string
	Role     int
}

const (
	tableName = "gd_user"
)

func QueryUserById(id int) *User {
	user := User{}

	sql, args, err := Select("*").From(tableName).Where(Eq{"user_id": id}).ToSQL()
	if err != nil {
		log.Log.Errorln("compose sql to query user by id failed")
		panic(err.Error())
	}

	result, err := Engine.Exec(sql, args)
	if err != nil {
		log.Log.Errorln("execute sql of  querying user by id failed")
		return nil
	}

	return &user
}

func QueryAllUserById(ids ...int) []*User {
	users := make([]*User, 0)

	return users
}
