package event

import (
	"helloworld/entities/user"
)

type Event struct {
	Id          int `gorm:"primarykey;autoIncrement"`
	User        user.User
	UserUid     string
	Picture     string
	Title       string
	Description string
}
