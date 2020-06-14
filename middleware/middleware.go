package middleware

import (
	"net/http"
	"encoding/json"
	"encoding/base64"
	"strings"
)

// CallIndex 
func ServiceClientMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		authorziation, ok := r.Header["X-Log-Autorization"]

		if !ok {
			resp := map[string]string{
				"message": "Authorization key is not set",
			}
			json.NewEncoder(w).Encode(resp)
		}


		b64, _ := base64.StdEncoding.DecodeString(authorziation[0])
		
		userID, userSecret := "jl_76h7y","8I67bsHjnYU8oi"
		
		secrets := strings.Split(string(b64), ":")

		if userID == secrets[0] && userSecret == secrets[1] {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

	})
}