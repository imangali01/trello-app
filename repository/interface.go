package repository

import (
	"github.com/Imangali2002/trello-app/view"
)

type UserInterface interface {
	RegisterUser(user view.UserCreate) error
	GetAllUsers() ([]view.UserInfo, error)
	GetUserByID(userID uint) (view.UserInfo, error)
}
