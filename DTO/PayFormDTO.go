package DTO

type PayForm struct {
	PayForm struct{
		Token string `json:"token"`
		FormUrl string `json:"form_url"`
	} `json:"pay_form"`
}

