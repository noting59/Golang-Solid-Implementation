package DTO

import "github.com/noting59/Golang-Solid-Implementation/models"

type CartDTO struct {
	Id int
	User *models.UserModel
	Product *models.ProductModel
}
