package handlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"okky/go-design-pattern/Repository/repositories"
)

/*
	- UserController is handler for user
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

type UserController struct {
	DB *gorm.DB
	repository repositories.UserRepository
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{
		DB:         db,
		repository: repositories.NewUserRepository(db),
	}
}

func (ctrl *UserController) GetUserByID() {
	id := 1
	user, err := ctrl.repository.Get(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}
