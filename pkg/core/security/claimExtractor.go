package security

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	ClaimsContextKey = "claims"
)

func GetCurrentUsername(ctx context.Context) (string, error) {
	username, ok := ExtractClaim(ctx, "username").(string)
	if ok {
		return username, nil
	}
	return "", errors.New("Error passing username")
}

func ExtractClaim(ctx context.Context, claimKey string) interface{} {
	claims := ctx.Value(ClaimsContextKey).(jwt.MapClaims)
	return claims[claimKey]
}
