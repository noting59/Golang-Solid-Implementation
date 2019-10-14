package interfaces

import (
	"test-task/DTO"
	"test-task/models"
)

type IProductService interface {
	GetProduct(userId int, productId int) (DTO.ProductDTO, error)
	GetById(id int) (models.ProductModel, error)
}
