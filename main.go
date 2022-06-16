package main

import (
	"fmt"
	"helloworld/config"
	user_event "helloworld/entities"
	"helloworld/entities/event"
	"helloworld/entities/user"
	"helloworld/middleware"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// gcp V&qTmCt:kOB)"T9`
func main() {
	appPort := getenv("PORT", "3000")

	routerR, err := setupRouter()

	err = routerR.Run(":" + appPort)
	if err != nil {
		return
	}

}

func setupRouter() (*gin.Engine, error) {
	var (
		dbUser         = getenv("DB_USER", "iustin")                 // e.g. 'my-db-user'
		dbPwd          = getenv("DB_PASS", "iustin")                 // e.g. 'my-db-password'
		unixSocketPath = getenv("INSTANCE_UNIX_SOCKET", "127.0.0.1") // e.g. '/cloudsql/project:region:instance'
		dbName         = getenv("DB_NAME", "sip")                    // e.g. 'my-database'
		dbPort         = getenv("DB_PORT", "5432")
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s TimeZone=Europe/Bucharest",
		dbUser, dbPwd, dbName, unixSocketPath, dbPort)

	dbPool, err := sql.Open("pgx", dbURI)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbPool,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		print(err)
		return nil, err
	}
	err = db.AutoMigrate(&event.Event{})
	if err != nil {
		print(err)
		return nil, err
	}
	err = db.AutoMigrate(&user_event.UserEvent{})
	if err != nil {
		print(err)
		return nil, err
	}
	// initialize gin Engine
	router := gin.Default()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	router.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	router.GET("/testPing", func(context *gin.Context) {
		value, _ := context.Get("UUID")
		print(value)
		context.JSON(http.StatusOK, gin.H{"data": value})
	})

	router.Use(middleware.AuthMiddleware)

	user.InitializeRoutes(router, db)
	event.InitializeRoutes(router, db)

	router.GET("/helloAuth", func(context *gin.Context) {
		value, _ := context.Get("UUID")
		print(value)
		context.JSON(http.StatusOK, gin.H{"data": value})
	})
	return router, nil
}
