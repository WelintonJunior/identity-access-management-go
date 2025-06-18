package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	repository "github.com/WelintonJunior/identity-access-management-go/repositories"
	"github.com/WelintonJunior/identity-access-management-go/services"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrMalformedRequest = "Invalid request payload"
)

// @Summary      User login
// @Description  Generates access and refresh tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body types.User true "Credentials"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/auth/login [post]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req types.User

		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   ErrMalformedRequest,
			})
		}

		accessToken, refreshToken, err := services.VerifyUser(req)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Email or password is incorrect",
			})
		}

		c.Locals("email", req.Email)

		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			HTTPOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":      true,
			"access_token": accessToken,
		})
	}
}

// @Summary      User registration
// @Description  Creates a new user in the database
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      types.User  true  "User data"
// @Success      201      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/auth/register [post]
func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req types.UserRegisterRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   ErrMalformedRequest,
			})
		}

		if req.Password != req.RepeatPassword {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Passwords do not match",
			})
		}

		user := types.User{
			FullName: req.FullName,
			Email:    req.Email,
			Password: req.Password,
			IsActive: true,
		}

		userRepo := repository.NewUserRepository()
		if err := userRepo.CreateUser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("failed to create user: %v", err),
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "User created successfully",
			"success": true,
		})
	}
}

// @Summary      Refresh access token
// @Description  Generates new access token using the refresh token in HttpOnly cookie
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/auth/refresh [get]
func RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies("refresh_token", "")
		if refreshToken == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Missing refresh token",
			})
		}

		email, err := auth.VerifyToken(refreshToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid refresh token",
			})
		}

		tokenRepo := repository.NewAuthTokenRepository()
		storedToken, err := tokenRepo.GetRefreshToken(c.Context(), "refresh-user:"+email)
		if err != nil || storedToken != refreshToken {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid or expired refresh token",
			})
		}

		newAccessToken, newRefreshToken, err := auth.GenerateTokens(email)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to generate tokens: %v", err),
			})
		}

		_ = tokenRepo.SetRefreshToken(c.Context(), "refresh-user:"+email, newRefreshToken, 7*24*time.Hour)

		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    newRefreshToken,
			HTTPOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"access_token": newAccessToken,
			"success":      true,
		})
	}
}
