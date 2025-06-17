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
	ErrMalFormedRequest = "Dados inválidos"
)

// @Summary 	Login do usuário
// @Description Gera access e refresh token
// @Tags 		Auth
// @Accept 		json
// @Produce 	json
// @Param 		request body types.User true "Credenciais"
// @Success 	200 {object} map[string]string
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/api/v1/auth/login [post]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request types.User

		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   ErrMalFormedRequest,
			})
		}

		accessToken, refreshToken, err := services.VerifyUser(request)

		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Email ou senha não conferem",
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			HTTPOnly: true,
			Secure:   false,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":      true,
			"access_token": accessToken,
		})
	}
}

// @Summary      Cadastro do usuário
// @Description  Cria um novo registro de usuário no banco de dados
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      types.User  true  "Dados do usuário"
// @Success      201      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/auth/register [post]
func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req types.User

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "invalid request body",
			})
		}

		user := types.User{
			FullName: req.FullName,
			Email:    req.Email,
			Password: req.Password,
			IsActive: true,
		}

		userRepository := repository.NewUserRepository()

		if err := userRepository.CreateUser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "user created",
			"success": true,
		})
	}
}

// @Summary      Atualiza o access token
// @Description  Gera novo access token com base no refresh token armazenado no cookie HttpOnly
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
				"error":   "token ausente",
			})
		}

		email, err := auth.VerifyToken(refreshToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "token inválido",
			})
		}

		authTokenRepository := repository.NewAuthTokenRepository()
		storedToken, err := authTokenRepository.GetRefreshToken(c.Context(), "refresh-user:"+email)
		if err != nil || storedToken != refreshToken {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "token inválido ou expirado",
			})
		}

		newAccessToken, newRefreshToken, err := auth.GenerateTokens(email)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("failed to generate token, %v", err),
			})
		}

		_ = authTokenRepository.SetRefreshToken(c.Context(), "refresh-user:"+email, newRefreshToken, 7*24*time.Hour)

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
