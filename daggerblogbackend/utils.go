package main

import (
	"crypto/tls"
	"daggerblogbackend/config"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func newRedisClient() *redis.Client {
	env := getEnv()
	var redisServer string
	var redisAuth string
	if env == "dev" {
		redisServer = config.REDIS_SERVER_DEV
		redisAuth = config.REDIS_AUTH_DEV
	} else if env == "qa" {
		redisServer = config.REDIS_SERVER_QA
		redisAuth = config.REDIS_AUTH_QA
	} else if env == "prod" {
		redisServer = config.REDIS_SERVER_PROD
		redisAuth = config.REDIS_AUTH_PROD
	} else {
		redisServer = config.REDIS_SERVER_DEV
		redisAuth = config.REDIS_AUTH_DEV
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + config.REDIS_PORT,
		Password: redisAuth,
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic("Error connecting redis")
	}

	return client
}

func setCacheItemVal(client *redis.Client, key string, val string, timeinsecs int) bool {
	err := client.Set(key, val, time.Duration(timeinsecs)*time.Second).Err()
	if err != nil {
		return false
	}
	return true
}

func deleteCacheItemByKey(client *redis.Client, key string) bool {
	client.Del(key)
	return true
}

func getCacheByKey(client *redis.Client, key string) string {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}
