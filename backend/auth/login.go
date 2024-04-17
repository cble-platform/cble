package auth

import (
	"net/http"
	"time"

	"github.com/cble-platform/cble/backend/config"
	"github.com/cble-platform/cble/backend/ent"
	"github.com/cble-platform/cble/backend/ent/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(cbleConfig *config.Config, client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username and password from the request body
		var input UserLoginInput
		if err := c.ShouldBind(&input); err != nil {
			DeleteAuthCookie(c, cbleConfig)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username/password"})
			return
		}

		// Lookup the user via username
		entUser, err := client.User.Query().Where(
			user.UsernameEQ(input.Username),
		).Only(c)
		if err != nil {
			DeleteAuthCookie(c, cbleConfig)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username/password"})
			return
		}

		// Compare the stored hashed password with input password hash
		if err = bcrypt.CompareHashAndPassword([]byte(entUser.Password), []byte(input.Password)); err != nil {
			DeleteAuthCookie(c, cbleConfig)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username/password"})
			return
		}

		// Generate the JWT token
		issuedAt := time.Now()
		expiresAt := issuedAt.Add(time.Minute * time.Duration(cbleConfig.Auth.SessionTimeout))
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer:    "CBLE",
			Subject:   entUser.Username,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(issuedAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ID:        uuid.NewString(),
		})

		// Sign the JWT token with the JWT key
		tokenString, err := jwtToken.SignedString([]byte(cbleConfig.Auth.JWTKey))
		if err != nil {
			DeleteAuthCookie(c, cbleConfig)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username/password"})
			return
		}

		// Set the JWT token as the session-token
		c.SetCookie("session-token", tokenString, int(time.Until(expiresAt).Seconds()), "/", cbleConfig.Server.Hostname, cbleConfig.Server.SSL, true)

		c.JSON(http.StatusOK, entUser)
	}
}
