package user_services

import (
	m "github.com/JrFerraz/GolangMongodbCRUD/models"
	userRepository "github.com/JrFerraz/GolangMongodbCRUD/repositories/user.repository"
)

//lo de retornar el error es simplemente buena practica
func Create(user m.User) error {
	err := userRepository.Create(user)
	return findErrors(err)
}

func Read() (m.Users, error) {

	users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user m.User, userId string) error {

	err := userRepository.Update(user, userId)

	return findErrors(err)
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)

	return findErrors(err)
}

func findErrors(err error) error {
	if err != nil {
		return err
	}
	return nil
}
