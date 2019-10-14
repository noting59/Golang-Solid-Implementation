package interfaces

import "github.com/noting59/Golang-Solid-Implementation/models"

type IOrderRepository interface {
	Create(productId int, userId int, name string, price float32) (int, error)
	GetById(orderId int) (models.OrderModel, error)
	Update(model models.OrderModel) (bool, error)
	List(userId int) []models.OrderModel
	GetUnProcessedOrders(status string) []models.OrderModel
}
