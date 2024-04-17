package auth

import (
	"net/http"

	"github.com/cble-platform/cble-backend/config"
	"github.com/gin-gonic/gin"
)

func DeleteAuthCookie(c *gin.Context, cbleConfig *config.Config) {
	c.SetCookie("session-token", "", -1, "/", cbleConfig.Server.Hostname, cbleConfig.Server.SSL, true)
}

func Logout(cbleConfig *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		DeleteAuthCookie(c, cbleConfig)
		c.JSON(http.StatusNoContent, gin.H{"message": "Successfully logged out user"})
	}
}
