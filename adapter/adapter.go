package adapter

import (
	"github.com/Nishad4140/user_service/entities"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{
		DB: db,
	}
}

func (u *UserAdapter) UserSignUp(req entities.User) (entities.User, error) {
	var res entities.User

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email"

	return res, u.DB.Raw(query, req.Name, req.Email, req.Password).Scan(&res).Error
}
