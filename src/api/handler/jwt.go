package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/rand"

	"github.com/linn-phyo/go_gin_clean_architecture/src/config"
)

type ConfigData struct {
	Config config.Config
}

func (cfg *ConfigData) GenerateToken(c *gin.Context) {
	// implement login logic here
	// user := c.PostForm("user")
	// pass := c.PostForm("pass")

	// // Throws Unauthorized error
	// if user != "john" || pass != "lark" {
	// 	return c.AbortWithStatus(http.StatusUnauthorized)
	// }

	// Create the Claims
	// claims := jwt.MapClaims{
	// 	"name":  "John Lark",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	// Create token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	username, password, ok := c.Request.BasicAuth()

	if ok {
		if FindCredential(username, password) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * 90).Unix(),
			})

			// fmt.Printf("\nGenerate JWT Secret Key >> %s", cfg.Config.JwtSecretKey)
			ss, err := token.SignedString([]byte(cfg.Config.JwtSecretKey))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"token": ss,
			})
		}
	}
}

// secrectkey := RandomString(32)
// log.Printf("SecrectKey >> %s", secrectkey)
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

type CredentialItem_Struct struct {
	ID          string `json:"id"`
	ConsumerKey string `json:"consumerKey"`
	KeyId       string `json:"keyId"`
	KeySecret   string `json:"keySecret"`
	IsActive    bool   `json:"isActive"`
}

type Credential_Struct struct {
	Credential_Item_Struct []CredentialItem_Struct `json:"credentials"`
}

func FindCredential(consumerkey string, secretkey string) bool {

	absCredentialsFilePath, _ := filepath.Abs("./src/data/credentials.json")

	// Open our jsonFile
	jsonFile, err := os.Open(absCredentialsFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened credentials.json")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := Credential_Struct{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	for _, item := range data.Credential_Item_Struct {
		fmt.Printf("ConsumerKey: %s, SecretKey: %s \n", item.ConsumerKey, item.KeySecret)
		if item.ConsumerKey == consumerkey && item.KeySecret == secretkey {
			fmt.Printf("Found a ConsumerKey: %s", item.ConsumerKey)
			return true
		}
	}

	fmt.Printf("Not  a ConsumerKey")

	return false
}
