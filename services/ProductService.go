package services

import (
	"test-task/DTO"
	"test-task/interfaces"
	"test-task/models"
)

type ProductService struct {
	interfaces.IProductRepository
	interfaces.ICartRepository
}

func (service *ProductService) GetProduct(userId int, productId int) (DTO.ProductDTO, error)  {
	product, err := service.IProductRepository.GetById(productId)

	if err != nil {
		return DTO.ProductDTO{}, err
	}

	productDTO := DTO.ProductDTO{Name:product.Name, Id:product.Id, Price:product.Price, InCart:false}

	cart, err := service.ICartRepository.CartByUserId(userId)

	if err != nil {
		return DTO.ProductDTO{}, err
	}

	if cart.ProductId == product.Id {
		productDTO.InCart = true
	}

	return productDTO, nil
}

func (service *ProductService) GetById(id int) (models.ProductModel, error) {
	res, err := service.IProductRepository.GetById(id)

	if err != nil {
		return models.ProductModel{}, err
	}

	return res, nil
}