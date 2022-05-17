package models

type User struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
