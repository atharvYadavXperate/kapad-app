package db

import (
	"context"

	customerror "github.com/atharvYadavXperate/kapad-app/domain/errors"
	"github.com/atharvYadavXperate/kapad-app/schema/users"
)

var (
	userCollection = "users"
)

func (db *Database) CreateUser(user users.UserSchema) (users.UserSchema, error) {
	user.HashPassword()
	_, _, err := db.CreateWithCustomId(context.Background(), userCollection, user.UserName, user)
	if err != nil {
		return users.UserSchema{}, customerror.ErrUserCreationFailed
	}
	return user, nil
}
