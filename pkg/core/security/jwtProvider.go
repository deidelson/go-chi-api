package security

import (
	"fmt"
	"github.com/deidelson/go-chi-api/pkg/core/date"
	"github.com/deidelson/go-chi-api/pkg/core/env"
	"github.com/dgrijalva/jwt-go"
)

type JwtProvider interface {
	CreateToken(claims jwt.MapClaims) (string, error)
	GetJwtClaims(token string) (jwt.MapClaims, error)
}

type JwtProviderImpl struct {
	secret                 []byte
	expirationTimeinMinuts int
}

var (
	jwtProviderInstance JwtProvider
)

func GetJwtProviderInstance() JwtProvider {
	if jwtProviderInstance == nil {
		jwtProviderInstance = &JwtProviderImpl{
			secret:                 []byte(env.GetEnvOrDefault(tokenSecretKey, "secret-local")),
			expirationTimeinMinuts: 60,
		}
	}
	return jwtProviderInstance
}

func (jwtProvider *JwtProviderImpl) CreateToken(claims jwt.MapClaims) (string, error) {
	expiration := date.Now()
	expiration = date.AddMinuts(expiration, env.GetEnvOrDefaultAsInt(tokenExpirationKey, 60))

	claims["exp"] = date.TimeToString(expiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtProvider.secret)

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jwtProvider *JwtProviderImpl) GetJwtClaims(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("security.token.signing.method")
		}
		return jwtProvider.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("security.token.signing.method")
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {

		exp, _ := date.StringToTime(claims["exp"].(string))

		if date.IsBeforeNow(exp) {
			return nil, fmt.Errorf("security.token.expired")
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("security.token.invalid")
	}
}
