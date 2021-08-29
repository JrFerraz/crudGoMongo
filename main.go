package main

import (
	"time"

	m "github.com/JrFerraz/GolangMongodbCRUD/models"
	userServices "github.com/JrFerraz/GolangMongodbCRUD/services/user.services"
)

func main() {

	user := m.User{
		Name:      "Noelia",
		Email:     "noeliavalero@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userServices.Create(user)

	userServices.Delete("612b88386636704135dba89c")
}
