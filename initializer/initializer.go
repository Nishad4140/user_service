package initializer

import (
	"github.com/Nishad4140/user_service/adapter"
	"github.com/Nishad4140/user_service/service"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) *service.UserService {
	adapter := adapter.NewUserAdapter(db)
	service := service.NewUserService(adapter)
	return service
}
