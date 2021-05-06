package public

import (
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/deidelson/go-chi-api/pkg/core/security"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var (
	loginHandlerInstance *handler
)

type handler struct {
	jwtProvider security.JwtProvider
}

func GetLoginHandlerInstance() routing.ApiHandler {
	if loginHandlerInstance == nil {
		loginHandlerInstance = &handler{
			jwtProvider: security.GetJwtProviderInstance(),
		}
	}
	return loginHandlerInstance
}

//Creates a fake jwt with 5 mins expiration
func (this *handler) generateFakeToken(w http.ResponseWriter, r *http.Request) {
	token, _ := this.jwtProvider.CreateToken(jwt.MapClaims{
		"username": "fakeUser",
	})
	web.Ok(w, token)
}

func (this *handler) GetBasePath() string {
	return "/public"
}

func (this *handler) GetMiddlewares() routing.Middlewares {
	return routing.Middlewares{}
}

func (this *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute{
		{
			Endpoint: "/token",
			Method:   "GET",
			Handler:  this.generateFakeToken,
		},
	}
}
