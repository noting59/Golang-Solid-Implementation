package interfaces

import "github.com/noting59/Golang-Solid-Implementation/models"

type IUserService interface {
	GetById(id int) (models.UserModel, error)
	Update(user models.UserModel) (bool, error)
}