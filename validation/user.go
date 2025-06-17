package validation

import (
	"errors"

	"github.com/WelintonJunior/identity-access-management-go/types"
)

func ValidateUser(user types.User) error {
	if user.Email == "" {
		return errors.New("Email não informado")
	}
	if user.FullName == "" {
		return errors.New("Nome não informado")
	}
	if user.Password == "" {
		return errors.New("Senha não informada")
	}
	return nil
}
