package db

type User struct {
	UserId   int    `xorm:"pk not null unique autoincr"`
	Password string `xorm:"not null"`
	UserName string `xorm:"not null"`
	Email    string `xorm:"null"`
	Role     int    `xorm:"not null"`
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
