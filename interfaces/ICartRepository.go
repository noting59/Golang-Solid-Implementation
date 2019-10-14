package interfaces

import (
	"test-task/models"
)

type ICartRepository interface {
	CartByUserId(userId int) (models.CartModel, error)
	AddProductToCart(userId int, productId int) (bool, error)
	DeleteByUserId(userId int) (bool, error)
}
