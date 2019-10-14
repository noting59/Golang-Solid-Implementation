package services

import (
	"test-task/interfaces"
	"test-task/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) GetById(orderId int) (models.UserModel, error) {
	res, err := service.IUserRepository.GetById(orderId)

	if err != nil {
		return  models.UserModel{}, err
	}

	return res, nil
}

func (service *UserService) Update(model models.UserModel) (bool, error)  {
	res, err := service.IUserRepository.Update(model)

	if err != nil {
		return  false, err
	}

	return res, nil
}