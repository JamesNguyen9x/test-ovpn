// Api handlers.
package handlers

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/JamesNguyen9x/test-ovpn/service/auth"
)

// Recover panics
func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"error": errors.New(fmt.Sprintf("%s", r)),
			}).Error("handlers: Handler panic")
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	c.Next()
}

// Log errors
func Errors(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("handlers: Handler error")
	}
}

// Auth requests
func Auth(c *gin.Context) {
	if c.Request.Header.Get("Origin") != "" ||
		c.Request.Header.Get("Referer") != "" ||
		c.Request.Header.Get("User-Agent") != "fvpn" ||
		subtle.ConstantTimeCompare(
			[]byte(c.Request.Header.Get("Auth-Key")),
			[]byte(auth.Key)) != 1 {

		c.AbortWithStatus(401)
		return
	}
	c.Next()
}

func Register(engine *gin.Engine) {
	engine.Use(Auth)
	engine.Use(Recovery)
	engine.Use(Errors)

	engine.GET("/events", eventsGet)
	engine.GET("/profile", profileGet)
	engine.POST("/profile", profilePost)
	engine.DELETE("/profile", profileDel)
	engine.PUT("/token", tokenPut)
	engine.DELETE("/token", tokenDelete)
	engine.GET("/ping", pingGet)
	engine.POST("/stop", stopPost)
	engine.POST("/restart", restartPost)
	engine.GET("/status", statusGet)
	engine.GET("/state", stateGet)
	engine.POST("/wakeup", wakeupPost)
}
