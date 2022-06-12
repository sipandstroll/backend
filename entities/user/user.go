package user

type User struct {
	Uid            string `gorm:"primarykey;unique;not null"`
	Name           string
	Email          string
	ProfilePicture string
	Age            uint
}
