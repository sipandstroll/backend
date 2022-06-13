package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitializeRoutes(engine *gin.Engine, db *gorm.DB) {

	engine.POST("/user", func(context *gin.Context) {
		uidN, exists := context.Get("UUID")

		uid, ok := uidN.(string)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if !exists {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		var newUser User
		if err := context.BindJSON(&newUser); err != nil {
			return
		}
		// the user must be authenticated to create its own row in users table
		if uid != newUser.Uid {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		tx := db.Create(newUser)
		if tx.Error == nil {
			context.JSON(http.StatusOK, newUser)
			return
		} else {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
	})

	engine.PUT("/user", func(context *gin.Context) {
		uidN, exists := context.Get("UUID")

		uid, ok := uidN.(string)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if !exists {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		var newUser User
		if err := context.BindJSON(&newUser); err != nil {
			return
		}
		// the user must be authenticated to update its own row in users table
		if uid != newUser.Uid {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		tx := db.Save(newUser)
		if tx.Error == nil {
			context.JSON(http.StatusOK, newUser)
			return
		} else {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
	})

	engine.GET("/user", func(context *gin.Context) {
		uidN, exists := context.Get("UUID")

		uid, ok := uidN.(string)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if !exists {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		var user User

		tx := db.Where(&User{Uid: uid}).First(&user)
		if tx.Error == nil {
			context.JSON(http.StatusOK, user)
			return
		} else {
			print(tx.Error)
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
	})
}
