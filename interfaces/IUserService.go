package interfaces

import "test-task/models"

type IUserService interface {
	GetById(id int) (models.UserModel, error)
	Update(user models.UserModel) (bool, error)
}