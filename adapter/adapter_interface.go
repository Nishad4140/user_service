package adapter

import "github.com/Nishad4140/user_service/entities"

type UserInterface interface {
	UserSignUp(req entities.User) (entities.User, error)
}
