package store

import (
	"github.com/gymnasy55/polcraz-rest-api-go/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
