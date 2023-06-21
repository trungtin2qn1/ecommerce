package utils

import (
	"context"
	"ecommerce/api/errors_reason"

	"github.com/go-kratos/kratos/v2/errors"
	jwtKratos "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	IdUser int64 `json:"id_user"`
	jwt.RegisteredClaims
}

func GetIdUserByJwt(ctx context.Context) (int64, error) {
	token, ok := jwtKratos.FromContext(ctx)
	if !ok {
		return 0, errors.Unauthorized(errors_reason.ErrorReason_UNAUTHORIZED.String(), "invalid token info")
	}
	claim, ok := token.(*Claims)
	if !ok {
		return 0, errors.Unauthorized(errors_reason.ErrorReason_UNAUTHORIZED.String(), "invalid claim info")
	}
	return claim.IdUser, nil
}
