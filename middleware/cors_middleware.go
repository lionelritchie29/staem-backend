package middleware

import(
	"net/http"
	"context"
)

type corsContextKey struct {
	name string
}

var corsCtxKey = &corsContextKey{"cors"}


func CorsMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		ctx := context.WithValue(r.Context(), corsCtxKey, w)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetCorsCtxValue(ctx context.Context) http.ResponseWriter {
	tolol, _ := ctx.Value(corsCtxKey).(http.ResponseWriter)
	return tolol
}

