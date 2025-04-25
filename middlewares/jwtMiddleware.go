package middlewares

import (
	"mi-proyecto/utils"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Token required")
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			utils.RespondWithError(c, http.StatusExpectationFailed, "Must contain the Bearer prefix")
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := jwt.MapClaims{}
		secretKey, err := utils.GetJWTPassword()
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Internal server error")
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Set("claims", claims)
		if len(roles) > 0 {
			userRole, ok := claims["role"].(string)
			if !ok || !slices.Contains(roles, userRole) {
				utils.RespondWithError(c, http.StatusForbidden, "Unauthorized role")
				return
			}
		}

		c.Next()
	}
}
