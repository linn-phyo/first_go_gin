package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/linn-phyo/go_gin_clean_architecture/src/config"
)

type ConfigData struct {
	Config config.Config
}

func (cfg *ConfigData) AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	err := cfg.validateToken(token)
	if err != nil {
		// c.AbortWithStatus(http.StatusUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
}

func (cfg *ConfigData) validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// fmt.Printf("\nVerify JWT Secret Key >> %s", cfg.Config.JwtSecretKey)
		return []byte(cfg.Config.JwtSecretKey), nil
	})

	return err
}
