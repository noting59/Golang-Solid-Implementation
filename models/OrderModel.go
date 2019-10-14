package models

import "time"

type OrderModel struct {
	Id int
	Name string
	ProductId int
	Price float32
	UserId int
	Status string
	FormToken string
	CreatedAt time.Time
	UpdatedAt time.Time
}

