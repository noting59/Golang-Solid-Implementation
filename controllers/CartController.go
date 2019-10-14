package controllers

import (
	"encoding/json"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"strconv"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
)

type CartController struct {
	interfaces.ICartService
	*renderer.Render
}

func (controller *CartController) AddToCart(r http.ResponseWriter, req *http.Request) {
	res := struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Code int `json:"code"`
	}{Status:"ok", Message: "Added to cart", Code: 200}

	productId := req.FormValue("productId")

	i, _ := strconv.Atoi(productId)

	_, err := controller.ICartService.AddProductToCart(1, i)

	if err != nil {
		res.Status = "ko"
		res.Message = err.Error()
		res.Code = 400
	}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	r.WriteHeader(res.Code)
	r.Header().Set("Content-Type", "application/json")
	r.Write(js)

	return
}

func (controller *CartController) Cart(r http.ResponseWriter, req *http.Request) {
	dto, err := controller.ICartService.GetCart(1)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "cart", dto)
	if err != nil {
		log.Fatal(err)
	}
}