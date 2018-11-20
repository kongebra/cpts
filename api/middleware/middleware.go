package middleware

import (
	"log"
	"net/http"
)

/*
Logger is a middleware for the mux-router
 */
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s \"%s %s %s\" %d\n", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, r.ContentLength)

		next.ServeHTTP(w, r)
	})
}