package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token has expired")
	ErrInvalidSigningMethod  = errors.New("invalid signing method")
	ErrMissingOrInvalidEmail = errors.New("missing or invalid email in token claims")
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

// Token durations
const (
	accessTokenExpiry  = 2 * time.Hour
	refreshTokenExpiry = 7 * 24 * time.Hour
)

// GenerateTokens creates both access and refresh tokens for a given user email.
func GenerateTokens(email string) (accessToken string, refreshToken string, err error) {
	if email == "" {
		return "", "", errors.New("email is required to generate tokens")
	}

	// Access Token
	atClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(accessTokenExpiry).Unix(),
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err = accessTokenObj.SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh Token
	rtClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(refreshTokenExpiry).Unix(),
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err = refreshTokenObj.SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// VerifyToken parses and validates a JWT token and extracts the user's email.
func VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return secretKey, nil
	})

	if err != nil {
		// Distinguish expired token
		var ve *jwt.ValidationError
		if errors.As(err, &ve) && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return "", ErrTokenExpired
		}
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", ErrInvalidToken
	}

	// Extract email claim
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return "", ErrMissingOrInvalidEmail
	}

	return email, nil
}

func CheckHashPassword(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
