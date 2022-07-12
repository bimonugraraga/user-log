package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}
