package db

import "gorm.io/gorm"

type User struct {
	UserId   int64
	Username string
	Password string
	Nickname string
}

func NewUser(user_name, password string) *User {
	return &User{
		Username: user_name,
		Password: password,
	}
}

func GetUserInfoByUserName(user_name string) (*User, error) {
	user := new(User)
	err := DB.Table("t_user").Where("username=?", user_name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user_name, password string) error {
	u := NewUser(user_name, password)
	return DB.Table("t_user").Create(&u).Error
}

func JudgeUser(user_name string) bool {
	user := new(User)
	err := DB.Table("t_user").Where("username=?", user_name).First(&user).Error
	return err == gorm.ErrRecordNotFound
}
