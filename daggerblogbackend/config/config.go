/* Copyright Fengchang Xie & Huifeng Zhang, 2018 All Rights Reserved */
package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

/**
 * Created by Fengchang on 2018/5/30 12:53.
 */

const (
	// dev
	RDB_PASSWORD         = "???"
	REDIS_SERVER_DEV     = "127.0.0.1"
	RDB_USER_DEV         = "???"
	RDB_NAME_DEV         = "???"
	DB_HOST              = "127.0.0.1"
	REDIS_AUTH_DEV       = "???"
	BASE_URL_DEV         = "http://localhost:???"
	IMG_DIR_DEV          = "/tmp"

	// qa
	RDB_PASSWORD_QA     = "???"
	REDIS_SERVER_QA     = "127.0.0.1"
	RDB_USER_QA         = "???"
	RDB_NAME_QA         = "????"
	REDIS_AUTH_QA       = "???"
	IMG_DIR_QA          = "/tmp"

	//prod
	RDB_PASSWORD_PROD     = "???$09ty"
	REDIS_SERVER_PROD     = "127.0.0.1"
	RDB_USER_PROD         = "???"
	RDB_NAME_PROD         = "???"
	BASE_URL_PROD         = "http://www.???.cc"
	IMG_DIR_PROD          = "/???"

	// common
	DB_PORT        = "9080"
	GIN_PORT       = ":90???"
	REDIS_PORT     = "6379"
	DIYI_MAIL      = "???@163.com"
	SMTP_SERVER    = "smtp.163.com"
	SMTP_PORT      = "465"
	MAIL_AUTH_CODE = "???"

	ITEMS_PER_PAGE = 100

	MAILNAME = "???"
)

var REDIS_CONN *redis.Client
var RDB_CONN *gorm.DB
