package DTO

import "test-task/models"

type CartDTO struct {
	Id int
	User *models.UserModel
	Product *models.ProductModel
}
