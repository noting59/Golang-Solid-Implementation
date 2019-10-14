package repositories

import (
	"fmt"
	"log"
	"test-task/interfaces"
	"test-task/models"
)

type UserRepository struct {
	interfaces.IDbHandler
}

func (repository *UserRepository) GetById(id int) (models.UserModel, error)  {
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM test.user WHERE id = %d", id))

	if err != nil {
		return models.UserModel{}, err
	}

	var user models.UserModel

	row.Next()
	row.Scan(&user.Id, &user.Name, &user.Email, &user.CardToken)

	return user, nil
}

func (repository *UserRepository) Update(user models.UserModel) (bool, error)  {
	sqlStatement := fmt.Sprintf("UPDATE test.user SET name = '%s', email = '%s', cardToken = '%s' WHERE id = %d", user.Name, user.Email, user.CardToken, user.Id)

	_,err := repository.Execute(sqlStatement)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
