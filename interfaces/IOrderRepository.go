package interfaces

import "test-task/models"

type IOrderRepository interface {
	Create(productId int, userId int, name string, price float32) (int, error)
	GetById(orderId int) (models.OrderModel, error)
	Update(model models.OrderModel) (bool, error)
	List(userId int) []models.OrderModel
	GetUnProcessedOrders(status string) []models.OrderModel
}
