package v1

import (
	"ecommerce/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (uc *AuthUsecase) generateAccessTokens(idUser int64) (string, string, error) {
	claims := &utils.Claims{
		IdUser: idUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(uc.config.JwtKey))
	if err != nil {
		return "", "", err
	}

	claims = &utils.Claims{
		IdUser: idUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * 30 * time.Hour)),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(uc.config.JwtKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
