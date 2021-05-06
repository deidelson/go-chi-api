package public

import (
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/deidelson/go-chi-api/pkg/core/security"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var (
	loginControllerInstance *handler
)

type handler struct {
	jwtProvider security.JwtProvider
}

func GetLoginHandlerInstance() routing.ApiHandler {
	if loginControllerInstance == nil {
		loginControllerInstance = &handler{
			jwtProvider: security.GetJwtProviderInstance(),
		}
	}
	return loginControllerInstance
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

func (controller *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute{
		{
			Endpoint: "/token",
			Method:  "GET",
			Handler: controller.generateFakeToken,
		},
	}
}
