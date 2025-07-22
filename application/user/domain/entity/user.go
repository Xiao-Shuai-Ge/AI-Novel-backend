package entity

import (
	"Ai-Novel/common/model"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type User struct {
	ID           int64
	Username     string
	Avatar       string
	Email        string
	Password     string
	HashPassword string
}

func NewUser(email, password string) User {
	return User{
		Email:        email,
		Password:     password,
		Username:     "",
		Avatar:       "",
		HashPassword: "",
	}
}

func (u *User) EncryptPassword() (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.HashPassword = string(hashPassword)
	return
}

func (u *User) SetID(id int64) {
	u.ID = id
}

func (u *User) DefaultUsername() {
	u.Username = "用户" + strconv.FormatInt(u.ID, 10)
}

func (u *User) Transform() *model.User {
	return &model.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.HashPassword,
		Username: u.Username,
		Avatar:   u.Avatar,
	}
}

func Form(user *model.User) User {
	return User{
		ID:           user.ID,
		Email:        user.Email,
		Password:     "",
		HashPassword: user.Password,
		Username:     user.Username,
		Avatar:       user.Avatar,
	}
}
