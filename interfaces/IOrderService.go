package interfaces

import (
	"test-task/DTO"
	"test-task/models"
)

type IOrderService interface {
	ProcessNewOrder(cardDTO DTO.CartDTO) (DTO.OrderDTO, error)
	ListOrders(userId int) (DTO.OrdersDTO, error)
	GetById(orderId int) (models.OrderModel, error)
	Update(model models.OrderModel) (bool, error)
	GetUnProcessedOrders(status string) []models.OrderModel
}
