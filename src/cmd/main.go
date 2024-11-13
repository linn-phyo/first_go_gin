package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linn-phyo/go_gin_clean_architecture/src/api/handler"
	"github.com/linn-phyo/go_gin_clean_architecture/src/api/middleware"
	"github.com/linn-phyo/go_gin_clean_architecture/src/api/routes"
	config "github.com/linn-phyo/go_gin_clean_architecture/src/config"
	"github.com/linn-phyo/go_gin_clean_architecture/src/db"
)

type Data struct {
	Item []Item `json:"item"`
}

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method string `json:"method"`
}

func main() {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	gormDB, err := db.ConnectDatabase(config)
	if err != nil {
		return
	}

	// os.Setenv("FOO", "Hello World!")
	// fmt.Println("FOO:", os.Getenv("FOO"))

	// Request JWT
	token_handler := handler.ConfigData{
		Config: config,
	}
	engine.GET("/token", token_handler.GenerateToken)

	middleware_handler := middleware.ConfigData{
		Config: config,
	}
	publicRouter := engine.Group("/api", middleware_handler.AuthorizationMiddleware)

	routes.Setup(gormDB, publicRouter)

	engine.Run(config.ServerAddress)

}
