package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Code":   http.StatusBadRequest,
		"Status": "Bad Request",
		"Data":   err.Error(),
	})
}

func ErrorInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Code":   http.StatusInternalServerError,
		"Status": "Internal Server Error",
		"Data":   err.Error(),
	})
}

func ErrorNotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"Code":   http.StatusNotFound,
		"Status": "Not Found",
		"Data":   err.Error(),
	})
}
