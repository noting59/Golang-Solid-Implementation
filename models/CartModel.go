package models

import "time"

type CartModel struct {
	Id int
	ProductId int
	UserId int
	CreatedAt time.Time
}
