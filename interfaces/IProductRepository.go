package interfaces

import (
	"github.com/noting59/Golang-Solid-Implementation/models"
)

type IProductRepository interface {
	GetById(id int) (models.ProductModel, error)
}
