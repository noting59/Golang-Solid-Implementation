package services

import (
	"github.com/noting59/Golang-Solid-Implementation/DTO"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
)

type CartService struct {
	interfaces.ICartRepository
	interfaces.IProductService
	interfaces.IUserService
}

func (service *CartService) AddProductToCart(userId int, productId int) (bool, error)  {
	res, err := service.ICartRepository.AddProductToCart(userId, productId)

	if err != nil {
		return false, err
	}

	return res, nil
}

func (service *CartService) GetCart(userId int) (DTO.CartDTO, error) {
	res, err := service.ICartRepository.CartByUserId(userId)

	if err != nil {
		return  DTO.CartDTO{}, err
	}

	user, err := service.IUserService.GetById(res.UserId)

	cartDTO := DTO.CartDTO{Id:res.Id, User:&user}

	product, err := service.IProductService.GetById(res.ProductId)

	if err != nil {
		return  DTO.CartDTO{}, err
	}

	cartDTO.Product = &product

	return cartDTO, nil
}

func (service *CartService) DeleteFromCart(userId int) (bool, error)  {
	res, err := service.ICartRepository.DeleteByUserId(userId)

	if err != nil {
		return false, err
	}

	return res, nil
}