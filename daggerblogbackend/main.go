package main

import (
	"daggerblogbackend/config"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

func main() {

	gin.SetMode(gin.ReleaseMode)
	// Set the router as the default one provided by Gin
	router = gin.Default()

	store := cookie.NewStore([]byte("secret"))
	// store.Options(sessions.Options{
	// 	MaxAge: 8000,
	// })
	router.Use(sessions.Sessions("fcblogbacksession", store))

	orig := "http://localhost:9???"

	env := getEnv()

	if env == "prod" {
		orig = config.BASE_URL_PROD
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{orig},
		AllowMethods:     []string{"PUT", "DELETE", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == orig
		},
		MaxAge: 12 * time.Hour,
	}))

	//router.MaxMultipartMemory = 8 << 20
	// Initialize the routes
	initializeRoutes()

	config.REDIS_CONN = newRedisClient()
	defer config.REDIS_CONN.Close()

	rdb := initRdbConnection()
	config.RDB_CONN = rdb
	defer config.RDB_CONN.Close()

	// Start serving the application
	log.Info("Starting fcblog backend...")
	router.Run(config.GIN_PORT)
}
func getEnv() string {
	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		if argsWithProg[1] == "prod" {
			return "prod"
		} else if argsWithProg[1] == "qa" {
			return "qa"
		} else {
			return "dev"
		}
	}
	return "dev"
}
func initRdbConnection() *gorm.DB {
	env := getEnv()
	var rdb_user string
	var rdb_password string
	var rdb_name string
	if env == "qa" {
		rdb_user = config.RDB_USER_QA
		rdb_password = config.RDB_PASSWORD_QA
		rdb_name = config.RDB_NAME_QA
	} else if env == "prod" {
		rdb_user = config.RDB_USER_PROD
		rdb_password = config.RDB_PASSWORD_PROD
		rdb_name = config.RDB_NAME_PROD
	} else {
		rdb_user = config.RDB_USER_DEV
		rdb_password = config.RDB_PASSWORD
		rdb_name = config.RDB_NAME_DEV
	}
	db, err := gorm.Open("mysql", rdb_user+":"+rdb_password+"@/"+rdb_name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Could not connect to db")
	}

	return db
}
