package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(DB *gorm.DB, group *gin.RouterGroup) {
	// All Public APIs
	NewUsersRouter(DB, group)
}
