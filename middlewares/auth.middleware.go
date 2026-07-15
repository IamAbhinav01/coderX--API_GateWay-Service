package middlewares

import (
	"Coderx/utils/formatters"
	"Coderx/utils/session"
	"net/http"
)

func AuthMiddleware(sm *session.SessionManager) func(http.Handler) http.Handler{

	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie,err := r.Cookie("session_id")

			if err != nil{
				formatters.ErrorResponse(w,http.StatusUnauthorized,"Unauthorized - No session cookie",err)
				return 
			}

			sessionData,err:= sm.Store().Get(r.Context(),cookie.Value)

			if err != nil || sessionData == nil{
				formatters.ErrorResponse(w,http.StatusUnauthorized,"Unauthorized - Invalid or expired session",err)
				return 
			}

			r.Header.Set("X-user-id",sessionData["user_id"])

			next.ServeHTTP(w,r)
		})
	}

}