package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonOK(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func doneWithStatus(c *gin.Context, statusCode int) {
	c.Writer.WriteHeader(statusCode)
	c.Done()
}

func noContent(c *gin.Context) {
	doneWithStatus(c, http.StatusNoContent)
}

func notFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func badRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func internalServerError(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
