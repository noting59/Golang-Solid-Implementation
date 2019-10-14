package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/noting59/Golang-Solid-Implementation/interfaces"
)

type PayController struct {
	*renderer.Render
	interfaces.IPayHandler
}

func (controller *PayController) InitPayAction(r http.ResponseWriter, req *http.Request) {
	orderId := chi.URLParam(req, "orderId")

	i, _ := strconv.Atoi(orderId)

	payForm, err := controller.IPayHandler.InitPay(i)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "pay", payForm)
	if err != nil {
		log.Fatal(err)
	}

}

func (controller *PayController) OneClickPayAction (r http.ResponseWriter, req *http.Request) {
	orderId := chi.URLParam(req, "orderId")

	i, _ := strconv.Atoi(orderId)

	payForm, err := controller.IPayHandler.OneClickPay(i)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "one-click-pay", payForm)
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *PayController) RefundAction (r http.ResponseWriter, req *http.Request) {
	orderId := chi.URLParam(req, "orderId")

	i, _ := strconv.Atoi(orderId)

	payForm, err := controller.IPayHandler.Refund(i)

	if err != nil {
		err = controller.Render.View(r, http.StatusNotFound, "notfound", nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = controller.Render.View(r, http.StatusOK, "refund", payForm)
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *PayController) PayProcess (r http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	_, err = controller.IPayHandler.ProcessPayment(body)

	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(map[string]string{"status": "ok"})

	r.Header().Set("Content-Type", "application/json")
	r.Write(res)

	return
}
