package main

import (
	"fmt"
	"helloworld/config"
	"helloworld/entities/event"
	"helloworld/entities/user"
	"helloworld/middleware"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// gcp V&qTmCt:kOB)"T9`
func main() {
	var (
		dbUser         = mustGetenv("DB_USER")              // e.g. 'my-db-user'
		dbPwd          = mustGetenv("DB_PASS")              // e.g. 'my-db-password'
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
		dbName         = mustGetenv("DB_NAME")              // e.g. 'my-database'
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
