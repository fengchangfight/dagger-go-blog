package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page

			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

func AuthNotRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.Next()
		} else {
			// Continue down the chain to handler etc
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
