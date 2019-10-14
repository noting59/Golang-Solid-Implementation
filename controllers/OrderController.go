package controllers

import (
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"test-task/interfaces"
)

type OrderController struct {
	interfaces.ICartService
	*renderer.Render
	interfaces.IOrderService
}

func (controller *OrderController) NewOrder(r http.ResponseWriter, req *http.Request) {
	cartDTO, err := controller.ICartService.GetCart(1)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	dto, err := controller.IOrderService.ProcessNewOrder(cartDTO)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "new-order", dto)
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *OrderController) OrdersList (r http.ResponseWriter, req *http.Request) {
	list, err := controller.IOrderService.ListOrders(1)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "order-list", list.Orders)
	if err != nil {
		log.Fatal(err)
	}
}
