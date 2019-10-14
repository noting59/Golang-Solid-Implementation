package interfaces

import "test-task/models"

type IUserRepository interface {
	GetById(id int) (models.UserModel, error)
	Update(user models.UserModel) (bool, error)
}
