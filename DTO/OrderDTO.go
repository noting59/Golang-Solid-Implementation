package DTO

import "github.com/noting59/Golang-Solid-Implementation/models"

type OrderDTO struct {
	Id int
	User *models.UserModel
	Product *models.ProductModel
	Name string
	Status string
}