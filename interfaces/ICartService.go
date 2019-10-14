package interfaces

import "test-task/DTO"

type ICartService interface {
	AddProductToCart(userId int, productId int) (bool, error)
	GetCart(userId int) (DTO.CartDTO, error)
	DeleteFromCart(userId int) (bool, error)
}
