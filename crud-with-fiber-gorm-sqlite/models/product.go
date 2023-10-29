package models

import "time"

type Product struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	SerialNo string `json:"serialNo"`
	CreatedAt time.Time
}