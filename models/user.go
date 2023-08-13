package models

import "time"

type User struct {
	Id        string    `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	PpUrl     string    `gorm:"column:ppUrl"`
	CreatedAt time.Time `gorm:"column:createdAt"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
