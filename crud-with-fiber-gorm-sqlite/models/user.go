package models

import "time"

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedAt time.Time
}
