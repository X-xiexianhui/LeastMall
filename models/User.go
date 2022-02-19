package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	Id       int
	Phone    string `json:"phone"`
	Password string `json:"password"`
	AddTime  int
	LastIp   string
	Email    string
	Status   int
}

func (User) TableName() string {
	return "user"
}
