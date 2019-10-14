package main

import (
	"github.com/go-chi/chi"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	productController := ServiceContainer().InjectProductController()
	cartController := ServiceContainer().InjectCartController()
	orderController := ServiceContainer().InjectOrderController()
	payController := ServiceContainer().InjectPayController()

	r := chi.NewRouter()
	r.HandleFunc("/", productController.Index)


	r.HandleFunc("/add-to-cart", cartController.AddToCart)
	r.HandleFunc("/cart", cartController.Cart)


	r.HandleFunc("/order/new", orderController.NewOrder)
	r.HandleFunc("/order/list", orderController.OrdersList)

	r.HandleFunc("/pay/{orderId}", payController.InitPayAction)
	r.HandleFunc("/pay/process", payController.PayProcess)
	r.HandleFunc("/pay/one-click/{orderId}", payController.OneClickPayAction)
	r.HandleFunc("/pay/refund/{orderId}", payController.RefundAction)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

