package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.comroarc0/hackernews/internal/jwt"
	"github.comroarc0/hackernews/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)

			r = r.WithContext(context.WithValue(r.Context(), userCtxKey, &user))
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
