package middlewares

import (
	"net/http"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authKey := r.URL.Query().Get("auth_key")

		if authKey != "123456789" {
			http.Error(w, "Unauthorized..", http.StatusUnauthorized)
			return
		}

		// check if there is an auth_token
		// if no, go next
		// if yes check if it is valid one
		// if not valid, go next
		// if valid, get user_id from it, then fetch the user, and add it to a context.

		next.ServeHTTP(w, r)
	})
}
