package main

import (
	"daggerblogbackend/config"
	"daggerblogbackend/entity"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func whoami(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	var displayName string

	if len(user.Nickname) > 0 {
		displayName = user.Nickname
	} else {
		displayName = username.(string)
	}

	if username == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		c.String(http.StatusOK, displayName)
	}
}

func login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	// compare plain text this is for personal password setting, do not encrpyt
	matched := (user.Password == password)

	if matched {
		session.Set("user", username) // In real world usage you'd set this to the users ID
		session.Set("uid", user.ID)   // In real world usage you'd set this to the users ID
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
