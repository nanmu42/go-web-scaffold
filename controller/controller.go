package controller

import (
	"net/http"
	"telescope/cache"
	"telescope/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller is where http logic lives
type Controller struct {
	L   *zap.Logger
	D   *database.DB
	Red *cache.Red
}

// Hello says hello world,
// also stands for health check.
func (con *Controller) Hello(c *gin.Context) {
	ok(c, "hello world!")
}

// NotFound 404 not found handler
func (con *Controller) NotFound(c *gin.Context) {
	c.PureJSON(http.StatusNotFound, R{
		Code: http.StatusNotFound,
		Msg:  http.StatusText(http.StatusNotFound),
		Data: nil,
	})
}

// MethodNotAllowed 405 method not allowed handler
func (con *Controller) MethodNotAllowed(c *gin.Context) {
	c.PureJSON(http.StatusMethodNotAllowed, R{
		Code: http.StatusMethodNotAllowed,
		Msg:  http.StatusText(http.StatusMethodNotAllowed),
		Data: nil,
	})
}

func (con *Controller) RobotsTXT(c *gin.Context) {
	c.String(http.StatusOK, `User-agent: *
Disallow: /`)
}