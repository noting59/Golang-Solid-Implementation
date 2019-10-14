package interfaces

import (
	"github.com/noting59/Golang-Solid-Implementation/DTO"
	"github.com/noting59/Golang-Solid-Implementation/models"
)

type IProductService interface {
	GetProduct(userId int, productId int) (DTO.ProductDTO, error)
	GetById(id int) (models.ProductModel, error)
}
