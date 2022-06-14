package main

import (
	"fmt"
	"helloworld/config"
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
	var (
		dbUser         = getenv("DB_USER", "iustin")                 // e.g. 'my-db-user'
		dbPwd          = getenv("DB_PASS", "iustin")                 // e.g. 'my-db-password'
		unixSocketPath = getenv("INSTANCE_UNIX_SOCKET", "127.0.0.1") // e.g. '/cloudsql/project:region:instance'
		dbName         = getenv("DB_NAME", "sip")                    // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s sslmode=disable TimeZone=Europe/Bucharest",
		dbUser, dbPwd, dbName, unixSocketPath)

	dbPool, err := sql.Open("pgx", dbURI)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbPool,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		print(err)
		return
	}
	err = db.AutoMigrate(&event.Event{})
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
	event.InitializeRoutes(router, db)

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
