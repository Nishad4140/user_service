package adapter

import "github.com/Nishad4140/user_service/entities"

type UserInterface interface {
	UserSignUp(req entities.User) (entities.User, error)
	UserLogin(email string) (entities.User, error)
	AdminLogin(email string) (entities.Admin, error)
	SupAdminLogin(email string) (entities.SuAdmin, error)
	AddAdmin(req entities.Admin) (entities.Admin, error)
	GetAllAdmins() ([]entities.Admin, error)
	GetAllUsers() ([]entities.User, error)
}
