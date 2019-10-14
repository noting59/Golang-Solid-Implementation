package controllers

import (
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
)

type ProductController struct {
	interfaces.IProductService
	*renderer.Render
}

func (controller *ProductController) Index(r http.ResponseWriter, req *http.Request) {
	productDTO, err := controller.IProductService.GetProduct(1, 1)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "product", productDTO)
	if err != nil {
		log.Fatal(err)
	}
}