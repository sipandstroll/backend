package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		var newUser User
		if err := context.BindJSON(&newUser); err != nil {
			return
		}
		db.Save(newUser)
	})
}
