package db

import (
	orm "opencodes/pkg/database"
)

type User struct {
	Id        int    `json:"id"` //主键ID
	Name      string `json:"name"`
	Email     string `json:"email"`
	Salt      string `json:"salt"`
	Password  string `json:"password"` //密码
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func CreateUser(user *User) error {
	db := orm.Eloquent.Table("t_user").Create(user)
	if db.Error != nil {

	}
	return nil
}
