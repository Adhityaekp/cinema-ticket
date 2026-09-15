package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	Type   string    `json:"type"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	userID uuid.UUID,
	role string,
	secret string,
	expireMinutes int,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID: userID,
		Role:   role,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expireMinutes) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(
	userID uuid.UUID,
	role string,
	secret string,
	expireDays int,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID: userID,
		Role:   role,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, expireDays)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}

func ValidateToken(
	tokenString string,
	secret string,
) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt == nil ||
		claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
