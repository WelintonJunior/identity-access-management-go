package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	repository "github.com/WelintonJunior/identity-access-management-go/repositories"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

const (
	MaxFailedLoginAttempts = 5
	LockoutDuration        = 15 * time.Minute
)

func VerifyUser(c *fiber.Ctx, session types.User) (string, string, error) {
	userRepository := repository.NewUserRepository()
	loginAttemptRepository := repository.NewLoginAttemptRepository()

	user, err := userRepository.FindUserByEmail(session.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	loginAttempt, err := loginAttemptRepository.FindLoginAttemptsByUserID(user.ID)
	if err != nil {
		return "", "", err
	}

	if loginAttempt == nil {
		loginAttempt = &types.LoginAttempt{
			UserID:              user.ID,
			FailedLoginAttempts: 0,
		}
		if saveErr := loginAttemptRepository.Save(loginAttempt); saveErr != nil {
			return "", "", errors.New("failed to handle login attempt")
		}
	}

	// If account is locked
	if loginAttempt.LockoutExpiresAt != nil && loginAttempt.LockoutExpiresAt.After(time.Now()) {
		return "", "", errors.New("account locked due to too many failed login attempts; please try again later")
	}

	// Verify password
	_, err = auth.CheckHashPassword(session.Password, user.Password)
	if err != nil {
		loginAttempt.FailedLoginAttempts++

		if loginAttempt.FailedLoginAttempts >= MaxFailedLoginAttempts {
			lockUntil := time.Now().Add(LockoutDuration)
			loginAttempt.LockoutExpiresAt = &lockUntil
		}

		if saveErr := loginAttemptRepository.Save(loginAttempt); saveErr != nil {
			fmt.Println("Warning: failed to update login attempt record:", saveErr)
		}

		return "", "", errors.New("invalid email or password")
	}

	// Successful login – reset attempts
	loginAttempt.FailedLoginAttempts = 0
	loginAttempt.LockoutExpiresAt = nil
	if saveErr := loginAttemptRepository.Save(loginAttempt); saveErr != nil {
		fmt.Println("Warning: failed to reset login attempt record:", saveErr)
	}

	// Generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(session.Email)
	if err != nil {
		return "", "", err
	}

	authTokenRepository := repository.NewAuthTokenRepository()
	refreshKey := "refresh-user:" + session.Email
	if err := authTokenRepository.SetRefreshToken(context.Background(), refreshKey, refreshToken, 7*24*time.Hour); err != nil {
		return "", "", err
	}

	// Audit log
	if err := LogAction(c, user.ID, "User logged in"); err != nil {
		fmt.Println("Warning: failed to log action:", err)
	}

	return accessToken, refreshToken, nil
}
