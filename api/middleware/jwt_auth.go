package middleware

import (
	"company-micro/domain"
	"company-micro/util/tokenutil"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

var (
	ErrNoTokenFound = errors.New("No token found")
)

const (
	TokenType = "Bearer"
)

type Middleware func(next http.Handler) http.Handler

// TODO: Create singleton?
func JwtAuthenticator(secretKey string, userIdCtxKey string, ur domain.UserRepository) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken, err := getTokenFromHeader(r)

			if err != nil {
				log.Println("Bad token: ", err.Error())
				render.Render(w, r, domain.NewGeneralErrResponse(http.StatusUnauthorized, "Bad token"))
				return
			}

			userId, err := tokenutil.ExtractIDFromToken(authToken, secretKey)
			if err != nil {
				log.Println("Bad user id: ", err.Error())
				render.Render(w, r, domain.NewGeneralErrResponse(http.StatusUnauthorized, "Not authorized"))
				return
			}

			_, err = ur.GetByID(r.Context(), userId)
			if err != nil {
				log.Println("User does not exist: ", err.Error())
				render.Render(w, r, domain.NewGeneralErrResponse(http.StatusUnauthorized, "Not authorized. User does not exist: "))
				return
			}

			ctx := context.WithValue(r.Context(), userIdCtxKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// getTokenFromHeader retrieves token from request
//
// "Authorization" request header example: "Authorization: BEARER T".
func getTokenFromHeader(r *http.Request) (string, error) {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:], nil
	}

	return "", ErrNoTokenFound
}
