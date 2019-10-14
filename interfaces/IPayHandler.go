package interfaces

import "github.com/noting59/Golang-Solid-Implementation/DTO"

type IPayHandler interface {
	InitPay (orderId int) (DTO.PayForm, error)
	ProcessPayment (body []byte) (bool, error)
	OneClickPay (orderId int) (bool, error)
	Refund (orderId int) (bool, error)
}