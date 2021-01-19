package middleware

import (
	"github.com/lionelritchie29/learn-gql-go/model"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"net/http"
	"context"
)

type authContextKey struct {
	name string
}

var authCtx = &authContextKey{"user"}

func Auth() func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			userId, err := helpers.ParseToken(tokenStr)

			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			db := database.GetInstance()
			defer db.Close()

			var loggedUser model.User
			db.Find(&loggedUser, userId)

			ctx := context.WithValue(r.Context(), authCtx, &loggedUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}