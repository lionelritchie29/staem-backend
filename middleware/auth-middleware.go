package middleware

import (
	"context"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/models"
	"net/http"
)

type authContextKey struct {
	name string
}

var authCtx = &authContextKey{"user"}


func Auth(next http.Handler) http.Handler {
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

			var loggedUser models.UserAccount
			db.Find(&loggedUser, userId)

			ctx := context.WithValue(r.Context(), "loggedUser", &loggedUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
