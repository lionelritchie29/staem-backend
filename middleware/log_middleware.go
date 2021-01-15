package middleware

import(
	"log"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r*http.Request) {
		log.Println("Log Middleware executing, URL " + r.URL.Path + " was hitting by " + r.RemoteAddr)
		next.ServeHTTP(w,r)
	})
}
