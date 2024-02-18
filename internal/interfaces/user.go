package interfaces

import (
	"github.com/Victor90001/prod/internal/entity"
)

type UserRepository interface {
	Login(user entity.User) (int, error)
}
