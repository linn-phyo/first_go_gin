package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	di "github.com/linn-phyo/go_gin_clean_architecture/src/di"
)

func NewUsersRouter(DB *gorm.DB, group *gin.RouterGroup) {

	userHandler, diErr := di.InitializeUsersAPI(DB)
	if diErr != nil {
		log.Fatal("Di Error : ", diErr)
	}

	group.GET("users", userHandler.FindAll)
	group.GET("users/:id", userHandler.FindByID)
	group.POST("users", userHandler.Save)
	group.DELETE("users/:id", userHandler.Delete)
}
