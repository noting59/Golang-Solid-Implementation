package repositories

import (
	"fmt"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
	"github.com/noting59/Golang-Solid-Implementation/models"
)

type ProductRepository struct {
	interfaces.IDbHandler
}

func (repository *ProductRepository) GetById(id int) (models.ProductModel, error)  {
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM test.product WHERE id = %d", id))

	if err != nil {
		return models.ProductModel{}, err
	}

	var product models.ProductModel

	row.Next()
	row.Scan(&product.Id, &product.Name, &product.Price)

	return product, nil
}
