package middleware

import (
	"github.com/gin-gonic/gin"
	"ms_novel/response"
	"net/http"
)

func GlobalExceptionMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			response.CustomerException(
				"The system started the car.",
				c,
				http.StatusInternalServerError,
			)
			return
		}
	}()
	// 加载完 defer recover，继续后续接口调用
	c.Next()
}
