package interfaces

import (
	"github.com/noting59/Golang-Solid-Implementation/models"
)

type ICartRepository interface {
	CartByUserId(userId int) (models.CartModel, error)
	AddProductToCart(userId int, productId int) (bool, error)
	DeleteByUserId(userId int) (bool, error)
}
