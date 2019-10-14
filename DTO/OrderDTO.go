package DTO

import "test-task/models"

type OrderDTO struct {
	Id int
	User *models.UserModel
	Product *models.ProductModel
	Name string
	Status string
}