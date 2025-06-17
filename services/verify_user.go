package services

import (
	"context"
	"errors"
	"time"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	repository "github.com/WelintonJunior/identity-access-management-go/repositories"
	"github.com/WelintonJunior/identity-access-management-go/types"
)

func VerifyUser(session types.User) (string, string, error) {
	userRepository := repository.NewUserRepository()

	user, err := userRepository.FindUserByEmail(session.Email)

	if err != nil {
		return "", "", err
	}

	isCorrect, err := auth.CheckHashPassword(session.Password, user.Password)

	if err != nil {
		return "", "", err
	}

	if isCorrect {
		accessToken, refreshToken, err := auth.GenerateTokens(session.Email)
		if err != nil {
			return "", "", err
		}

		authTokenRepository := repository.NewAuthTokenRepository()
		key := "refresh-user:" + session.Email

		if err := authTokenRepository.SetRefreshToken(context.Background(), key, refreshToken, 7*24*time.Hour); err != nil {
			return "", "", err
		}

		return accessToken, refreshToken, nil
	} else {
		return "", "", errors.New("Incorrect password")
	}
}
