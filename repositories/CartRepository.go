package repositories

import (
	"fmt"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
	"github.com/noting59/Golang-Solid-Implementation/models"
	"time"
)

type CartRepository struct {
	interfaces.IDbHandler
}

func (repository *CartRepository) CartByUserId(userId int) (models.CartModel, error)  {
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM test.cart WHERE userId = %d", userId))

	if err != nil {
		return models.CartModel{}, err
	}

	var cart models.CartModel

	row.Next()
	row.Scan(&cart.Id, &cart.ProductId, &cart.UserId, &cart.CreatedAt)

	return cart, nil
}

func (repository *CartRepository) DeleteByUserId(userId int) (bool, error)  {
	_, err := repository.Query(fmt.Sprintf("DELETE FROM test.cart WHERE userId = %d", userId))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *CartRepository) AddProductToCart(userId int, productId int) (bool, error)  {
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	sqlStatement := fmt.Sprintf("INSERT INTO test.cart (productId, userId, createdAt) VALUES (%d, %d, '%s')", productId, userId, date)
	_, err := repository.Execute(sqlStatement)

	if err != nil {
		return false, err
	}
	return true, nil
}