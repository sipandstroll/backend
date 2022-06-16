package user_event

import (
	"helloworld/entities/event"
	"helloworld/entities/user"
)

type UserEvent struct {
	Id      int `gorm:"primarykey;autoIncrement"`
	User    user.User
	UserUid string
	Event   event.Event
	EventId int
	Answer  string // Go / Declined TODO: Refactor this to an enum in database
}
