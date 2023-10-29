package models

import "time"

type Order struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	ProductRefer int `json:"productId"`
	Product Product `gorm:"foreignKey:ProductRefer"`
	UserRefer int `json:"userId"`
	User User `gorm:"foreignKey:UserRefer"`
	CreatedAt time.Time
}