package db

type User struct {
	UserId   int    `xorm:"pk not null unique autoincr"`
	Password string `xorm:"not null"`
	UserName string `xorm:"not null"`
	Email    string `xorm:"null"`
	Role     int    `xorm:"not null"`
}

func (user *User) TableName() string {
	return TableUser
}

func QueryUserById(id int) (*User, error) {
	user, err := queryById(id, new(User))
	return user.(*User), err
}
