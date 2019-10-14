package interfaces

import "test-task/DTO"

type IPayHandler interface {
	InitPay (orderId int) (DTO.PayForm, error)
	ProcessPayment (body []byte) (bool, error)
	OneClickPay (orderId int) (bool, error)
	Refund (orderId int) (bool, error)
}