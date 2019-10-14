package interfaces

import (
	"test-task/models"
)

type IProductRepository interface {
	GetById(id int) (models.ProductModel, error)
}
