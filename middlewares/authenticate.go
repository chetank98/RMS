package middlewares

import (
	"RMS/database/dbHelper"
	"RMS/models"
	"RMS/utils"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

type ContextKeys string

const (
	userContext ContextKeys = "userContext"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, nil, "authorization header missing")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.RespondError(w, http.StatusUnauthorized, nil, "bearer token missing")
			return
		}

		token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method") // Invalid signing method error
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if parseErr != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, parseErr, "invalid token")
			return
		}

		claimValues, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, nil, "invalid token claims")
			return
		}

		sessionId := claimValues["sessionID"].(string)

		archivedAt, err := dbHelper.GetArchivedAt(sessionId)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if archivedAt != nil {
			utils.RespondError(w, http.StatusUnauthorized, nil, "invalid token")
			return
		}

		user := &models.UserCtx{
			UserId:    claimValues["userId"].(string),
			SessionId: sessionId,
			Role:      models.Role(claimValues["role"].(string)),
			Name:      claimValues["name"].(string),
			Email:     claimValues["email"].(string),
		}

		ctx := context.WithValue(r.Context(), userContext, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func UserContext(r *http.Request) *models.UserCtx {
	if user, ok := r.Context().Value(userContext).(*models.UserCtx); ok {
		return user
	}
	return nil
}
