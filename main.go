package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"helloworld/config"
	"helloworld/entities/user"
	"helloworld/middleware"
	"net/http"
)

// gcp V&qTmCt:kOB)"T9`
func main() {
	dsn := "host=127.0.0.1 user=iustin password=iustin dbname=sip port=5432 sslmode=disable TimeZone=Europe/Bucharest"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		print(err)
		return
	}

	print(db, err)

	// initialize gin Engine
	router := gin.Default()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	router.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	router.Use(middleware.AuthMiddleware)

	user.InitializeRoutes(router, db)

	router.GET("/helloAuth", func(context *gin.Context) {
		value, _ := context.Get("UUID")
		print(value)
		context.JSON(http.StatusOK, gin.H{"data": value})
	})

	err = router.Run(":3000")
	if err != nil {
		return
	}

	_, _ = router, firebaseAuth
	_ = db
}
