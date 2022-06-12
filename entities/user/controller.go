package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitializeRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.POST("/user", func(context *gin.Context) {
		var newUser User
		if err := context.BindJSON(&newUser); err != nil {
			return
		}
		db.Create(newUser)
	})

	engine.PUT("/user", func(context *gin.Context) {
		var newUser User
		if err := context.BindJSON(&newUser); err != nil {
			return
		}
		db.Save(newUser)
	})

	engine.GET("/user", func(context *gin.Context) {
		var user User
		if err := context.BindJSON(&user); err != nil {
			return
		}
		tx := db.First(&user)
		if tx.Error == nil {
			context.JSON(http.StatusOK, user)
		} else {
			context.JSON(http.StatusNotFound, gin.H{})
		}
	})
}
