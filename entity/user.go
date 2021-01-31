package entity

type User struct {
	UserId uint `gorm:"primaryKey"`
	Name   string
	Age    int
}

func (user *User) TableName() string {
	return "users"
}
