package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"zero-chat/common/ctxdata"

	"github.com/golang-jwt/jwt"
)

type JwtMiddleware struct {
	JwtSecret string
}

func NewJwtMiddleware(secret string) *JwtMiddleware {
	return &JwtMiddleware{
		JwtSecret: secret,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.URL.Query()
		tokenString := val.Get("token")
		if tokenString == "" {
			http.Error(w, "Authorization token not found", http.StatusUnauthorized)
			return
		}
		// 解析 JWT Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
			return
		}
		userId, ok := claims[string(ctxdata.CtxKeyJwtUserId)].(float64)
		if !ok {
			http.Error(w, "Failed to parse CtxKeyJwtUserId", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ctxdata.CtxKeyJwtUserId, json.Number(strconv.FormatFloat(userId, 'f', -1, 64)))

		next(w, r.WithContext(ctx))
	}
}
