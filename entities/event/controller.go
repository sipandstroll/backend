package event

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitializeRoutes(engine *gin.Engine, db *gorm.DB) {

	engine.GET("/event", func(context *gin.Context) {
		var events []Event
		tx := db.Find(&events)

		if tx.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
		context.JSON(http.StatusOK, events)
	})

	engine.GET("/event/mine", func(context *gin.Context) {
		var events []Event
		uidN, _ := context.Get("UUID")
		uid, _ := uidN.(string)

		tx := db.Where("user_uid = ?", uid).Find(&events)

		if tx.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
		context.JSON(http.StatusOK, events)
	})

	engine.GET("/event/:eventId", func(context *gin.Context) {
		eventId := context.Param("eventId")
		var event Event
		tx := db.First(&event, eventId)

		if tx.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
		context.JSON(http.StatusOK, event)
	})

	engine.POST("/event", func(context *gin.Context) {
		var event Event
		if err := context.BindJSON(&event); err != nil {
			print(err)
			context.JSON(http.StatusBadRequest, err)
			return
		}
		uidN, _ := context.Get("UUID")
		uid, _ := uidN.(string)

		if uid != event.UserUid {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		tx := db.Create(&event)
		if tx.Error != nil {
			context.JSON(http.StatusBadRequest, tx.Error)
			return
		}
		context.JSON(http.StatusOK, event)

	})

	engine.DELETE("/event/:eventId", func(context *gin.Context) {
		eventId := context.Param("eventId")
		var event Event
		tx := db.First(&event, eventId)

		if tx.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}

		uidN, _ := context.Get("UUID")
		uid, _ := uidN.(string)

		if event.UserUid != uid {
			context.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		txDeleted := db.Delete(event)
		if txDeleted.Error != nil {
			context.JSON(http.StatusBadRequest, txDeleted.Error)
			return
		}

		context.JSON(http.StatusOK, event)
	})

}
