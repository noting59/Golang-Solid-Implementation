package services

import (
	"fmt"
	"test-task/DTO"
	"test-task/interfaces"
	"test-task/models"
	"time"
)

type OrderService struct {
	interfaces.IOrderRepository
	interfaces.IProductService
	interfaces.IUserService
}

func (service *OrderService) ProcessNewOrder(cardDTO DTO.CartDTO) (DTO.OrderDTO, error)  {
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	name := fmt.Sprintf("%s_%d_%s", cardDTO.Product.Name, cardDTO.User.Id, date)

	newOrderId, err := service.IOrderRepository.Create(cardDTO.Product.Id, cardDTO.User.Id, name, cardDTO.Product.Price)

	if err != nil {
		return DTO.OrderDTO{}, err
	}

	newOrder, err := service.IOrderRepository.GetById(newOrderId)

	if err != nil {
		return DTO.OrderDTO{}, err
	}

	dto := DTO.OrderDTO{
		Id:newOrder.Id,
		User: cardDTO.User,
		Status: newOrder.Status,
		Product: cardDTO.Product,
		Name:newOrder.Name,
	}

	return dto, nil
}

func (service *OrderService) ListOrders(userId int) (DTO.OrdersDTO, error) {
	list := service.IOrderRepository.List(userId)

	user, err := service.IUserService.GetById(userId)

	if err != nil {
		return DTO.OrdersDTO{}, err
	}

	dto := DTO.OrdersDTO{}

	for _, element := range list {
		product, err := service.IProductService.GetById(element.ProductId)

		if err != nil {
			return DTO.OrdersDTO{}, err
		}

		orderDto := DTO.OrderDTO{
			Id: element.Id,
			User: &user,
			Status: element.Status,
			Product: &product,
			Name: element.Name,
		}

		dto.Orders = append(dto.Orders, orderDto)
	}

	return dto, nil
}

func (service *OrderService) GetById(orderId int) (models.OrderModel, error) {
	res, err := service.IOrderRepository.GetById(orderId)

	if err != nil {
		return  models.OrderModel{}, err
	}

	return res, nil
}

func (service *OrderService) Update(model models.OrderModel) (bool, error)  {
	res, err := service.IOrderRepository.Update(model)

	if err != nil {
		return  false, err
	}

	return res, nil
}

func (service *OrderService) GetUnProcessedOrders(status string) []models.OrderModel {
	list := service.IOrderRepository.GetUnProcessedOrders(status)

	return list
}