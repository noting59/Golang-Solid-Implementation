package repositories

import (
	"fmt"
	"log"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
	"github.com/noting59/Golang-Solid-Implementation/models"
	"time"
)

const StatusCreated = "created"
const StatusInited = "inited"

type OrderRepository struct {
	interfaces.IDbHandler
}

func (repository *OrderRepository) GetById(orderId int) (models.OrderModel, error) {
	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM test.order WHERE id = %d", orderId))

	if err != nil {
		return models.OrderModel{}, err
	}

	var order models.OrderModel

	row.Next()
	row.Scan(&order.Id, &order.Name, &order.ProductId, &order.Price, &order.Status, &order.UserId, &order.FormToken, &order.CreatedAt, &order.UpdatedAt)

	return order, nil
}

func (repository *OrderRepository) Update(model models.OrderModel) (bool, error)  {
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	sqlStatement := fmt.Sprintf("UPDATE test.order SET productId = %d, userId = %d, price = %f, status = '%s', name = '%s', formToken = '%s', updatedAt = '%s' WHERE id = %d", model.ProductId, model.UserId, model.Price, model.Status, model.Name, model.FormToken, date, model.Id)

	_,err := repository.Execute(sqlStatement)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func (repository *OrderRepository) Create(productId int, userId int, name string, price float32) (int, error)  {
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	sqlStatement := fmt.Sprintf("INSERT INTO test.order (productId, userId, createdAt, updatedAt, name, status, price) VALUES (%d, %d, '%s', '%s', '%s', '%s', %f) RETURNING id", productId, userId, date, date, name, StatusInited, price)
	id := 0
	err := repository.Connection().QueryRow(sqlStatement).Scan(&id)

	if err != nil {
		log.Fatal(err.Error())
	}

	return id, nil
}

func (repository *OrderRepository) List(userId int) []models.OrderModel {
	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM test.order WHERE userId = %d ORDER BY id", userId))

	if err != nil {
		log.Fatal(err.Error())
	}

	var orders []models.OrderModel

	for row.Next() {
		order := models.OrderModel{}
		row.Scan(&order.Id, &order.Name, &order.ProductId, &order.Price, &order.Status, &order.UserId, &order.FormToken, &order.CreatedAt, &order.UpdatedAt)

		orders = append(orders, order)
	}

	return orders
}

func (repository *OrderRepository) GetUnProcessedOrders(status string) []models.OrderModel {
	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM test.order WHERE status = '%s'", status))

	if err != nil {
		log.Fatal(err.Error())
	}

	var orders []models.OrderModel

	for row.Next() {
		order := models.OrderModel{}
		row.Scan(&order.Id, &order.Name, &order.ProductId, &order.Price, &order.Status, &order.UserId, &order.FormToken, &order.CreatedAt, &order.UpdatedAt)

		orders = append(orders, order)
	}

	return orders
}