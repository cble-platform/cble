package auth

import (
	"context"
	"net/http"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/cble-platform/cble-backend/internal/contexts"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Looks up the ENT user object from the session token and injects into context
func AuthMiddleware(cbleConfig *config.Config, client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session token from request cookies
		tokenString, err := c.Cookie("session-token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}

		// Unmarshal the JWT data from the session token
		claims := &jwt.RegisteredClaims{}
		jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cbleConfig.Auth.JWTKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}
		if !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}

		// Lookup the user based on the JWT token subject
		entUser, err := client.User.Query().Where(user.UsernameEQ(claims.Subject)).Only(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}

		// Inject the user object into request context
		ctx := context.WithValue(c.Request.Context(), contexts.USER_CTX_KEY, entUser)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
