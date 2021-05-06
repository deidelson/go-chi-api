package middleware

import (
	"context"
	"github.com/deidelson/go-chi-api/pkg/core/security"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"net/http"
)

var (
	ClaimsContextKey string = "claims"
)

type SecurityError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := security.GetJwtProviderInstance().GetJwtClaims(r.Header.Get("token"))

		if err != nil {
			errorBody := &SecurityError{
				ErrorCode: 401,
				Message:   err.Error(),
			}

			web.WriteSecurityError(w, errorBody)
		} else {
			requestWithContext := r.WithContext(context.WithValue(r.Context(), ClaimsContextKey, claims))
			next.ServeHTTP(w, requestWithContext)
		}
	})
}
