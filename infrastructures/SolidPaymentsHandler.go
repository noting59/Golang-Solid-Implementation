package infrastructures

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"test-task/DTO"
	"test-task/interfaces"
	"test-task/models"
	"test-task/repositories"
)

const ApiUrl = "https://pay.sp-stage.us/api/v1/"
const MerchantId = "unicorn"
const PrivateKey = "20c20ee3-4173-4daa-87e5-dbcce8c7949d"

type SolidPayments struct {
	interfaces.IOrderService
	interfaces.IUserService
	interfaces.ICartService
}

type InitPayData struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	CustomerEmail string `json:"customer_email"`
	IpAddress string `json:"ip_address"`
	OrderId string `json:"order_id"`
	OrderDescription string `json:"order_description"`
	Platform string `json:"platform"`
}

type OneClickPayData struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
	CustomerEmail string `json:"customer_email"`
	IpAddress string `json:"ip_address"`
	OrderId string `json:"order_id"`
	OrderDescription string `json:"order_description"`
	Platform string `json:"platform"`
	RecurringToken string `json:"recurring_token"`
	PaymentType string `json:"payment_type"`
}

type RefundData struct {
	Amount int `json:"amount"`
	OrderId string `json:"order_id"`
}

type StatusCheckData struct {
	OrderId string `json:"order_id"`
}

type ProcessPayData struct {
	Transactions map[string]struct{
		Id string `json:"id"`
		Status string `json:"status"`
		Operation string `json:"operation"`
		Card struct{
			CardToken struct{
				Token string `json:"token"`
			} `json:"card_token"`
		} `json:"card"`
	} `json:"transactions"`
	Order struct{
		OrderId string `json:"order_id"`
		Status string `json:"status"`
		Amount int `json:"amount"`
	} `json:"order"`
	Type string `json:"type"`
}

func (solid *SolidPayments) InitPay (orderId int) (DTO.PayForm, error) {
	order, err := solid.IOrderService.GetById(orderId)

	if err != nil {
		return DTO.PayForm{}, err
	}

	user, err := solid.IUserService.GetById(order.UserId)

	if order.FormToken != "" {
		payForm := DTO.PayForm{}
		payForm.PayForm.Token = order.FormToken

		return payForm, nil
	}

	if err != nil {
		return DTO.PayForm{}, err
	}

	payLoadJson, _ := json.Marshal(InitPayData{
		Amount: int(order.Price * (100)),
		Currency: "USD",
		CustomerEmail: user.Email,
		IpAddress: "8.8.8.8",
		OrderId: strconv.Itoa(order.Id),
		OrderDescription: order.Name,
		Platform: "WEB",
	})

	res := solid.makeRequest("init-payment", payLoadJson)

	log.Print(res.StatusCode)

	if res.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		payForm :=  DTO.PayForm{}
		err = decoder.Decode(&payForm)

		order.FormToken = payForm.PayForm.Token
		order.Status = repositories.StatusCreated

		_, err = solid.IOrderService.Update(order)

		if err != nil {
			log.Fatal(err)
		}

		return payForm, nil
	}

	return DTO.PayForm{}, nil
}

func (solid * SolidPayments) ProcessPayment (body []byte) (bool, error) {
	var processData ProcessPayData

	err := json.Unmarshal(body, &processData)

	if processData.Type == "orderStatus" {
		return false, err
	}

	i, _ := strconv.Atoi(processData.Order.OrderId)

	order, err := solid.IOrderService.GetById(i)

	if err != nil {
		return false, err
	}

	user, err := solid.IUserService.GetById(order.UserId)

	if err != nil {
		return false, err
	}

	order.Status = processData.Order.Status

	if order.Status == "approved" {
		_, err = solid.ICartService.DeleteFromCart(order.UserId)

		if err != nil {
			return false, err
		}
	}

	for _,elem := range processData.Transactions {
		if elem.Operation == "pay" && elem.Status == "success" {
			if elem.Card.CardToken.Token != "" {
				user.CardToken = elem.Card.CardToken.Token
			}
		}
	}

	_, err = solid.IOrderService.Update(order)

	if err != nil {
		return false, err
	}

	_, err = solid.IUserService.Update(user)

	return true, nil
}

func (solid *SolidPayments) OneClickPay (orderId int) (bool, error){
	order, err := solid.IOrderService.GetById(orderId)

	if err != nil {
		return false, err
	}

	user, err := solid.IUserService.GetById(order.UserId)

	if err != nil {
		return false, err
	}

	payLoadJson, _ := json.Marshal(OneClickPayData{
		Amount:           int(order.Price * (100)),
		Currency:         "USD",
		CustomerEmail:    user.Email,
		IpAddress:        "8.8.8.8",
		OrderId: strconv.Itoa(order.Id),
		OrderDescription: order.Name,
		Platform:         "WEB",
		RecurringToken:   user.CardToken,
		PaymentType:      "1-click",
	})

	res := solid.makeRequest("recurring", payLoadJson)

	if res.StatusCode == http.StatusOK {
		solid.processStatusOfOrder(order, res.Body, true)

		return true, nil
	}

	return false, nil
}

func (solid *SolidPayments) Refund (orderId int) (bool, error) {
	order, err := solid.IOrderService.GetById(orderId)

	if err != nil {
		return false, err
	}

	payLoadJson, _ := json.Marshal(RefundData{
		Amount:  int(order.Price * (100)),
		OrderId: strconv.Itoa(order.Id),
	})

	res := solid.makeRequest("refund", payLoadJson)

	if res.StatusCode == http.StatusOK {
		solid.processStatusOfOrder(order, res.Body, true)

		return true, nil
	}

	return false, nil
}

func (solid *SolidPayments) StatusCheck (orderId int) (bool, error) {
	order, err := solid.IOrderService.GetById(orderId)

	if err != nil {
		return false, err
	}

	payLoadJson, _ := json.Marshal(StatusCheckData{
		OrderId: strconv.Itoa(order.Id),
	})

	res := solid.makeRequest("status", payLoadJson)

	if res.StatusCode == http.StatusOK {
		solid.processStatusOfOrder(order, res.Body, false)

		return true, nil
	}

	return false, nil
}

func getSignature(input, key string) string {
	keyForSign := []byte(key)
	h := hmac.New(sha512.New, keyForSign)
	h.Write([]byte(input))

	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}

func (solid *SolidPayments) processStatusOfOrder (order models.OrderModel, body io.Reader, deleteFromCart bool) {
	decoder := json.NewDecoder(body)

	processData :=  ProcessPayData{}
	err := decoder.Decode(&processData)

	if order.Status == processData.Order.Status {
		return
	}

	order.Status = processData.Order.Status

	_, err = solid.IOrderService.Update(order)

	if deleteFromCart {
		_, err = solid.ICartService.DeleteFromCart(order.UserId)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func (solid *SolidPayments) makeRequest (url string, payloadJson []byte) *http.Response {
	payloadString := MerchantId + string(payloadJson) + MerchantId

	req, err := http.NewRequest("POST", ApiUrl+url, bytes.NewBuffer(payloadJson))
	req.Header.Set("Signature", getSignature(payloadString, PrivateKey))
	req.Header.Set("Merchant", MerchantId)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return res
}