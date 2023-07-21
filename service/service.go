package service

import (
	"context"
	"net/http"

	"recitas/config"
	"recitas/modules/recipes"
	"recitas/modules/users"
	"recitas/proto/rec"

	"github.com/twitchtv/twirp"
)

const cookieName = "user_token"

type Api struct {
	recipes recipes.IRecipes
	users   users.IUsers
}

func NewService(cfg config.Config) http.Handler {
	prefix := twirp.WithServerPathPrefix("/twirp")
	return http.Handler(
		rec.NewApiGatewayServer(
			&Api{
				recipes.NewRecipeModule(cfg),
				users.NewUserModule(cfg),
			},
			prefix,
		),
	)
}

func HeadersTransmitter(ctx context.Context, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("user_token")
		cookieVal := ""
		if cookie != nil {
			cookieVal = cookie.Value
		}
		ctx = context.WithValue(ctx, users.UserTokenKey, cookieVal)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}
