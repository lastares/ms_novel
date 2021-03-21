package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ms_novel/global"
	"ms_novel/response"
	"net/http"
)

func FormValidateException(c *gin.Context, err error) {
	// 写法1，类似php数组 $a[0]
	errorMessage := err.(validator.ValidationErrors)[0].Translate(global.Translator)
	c.JSON(http.StatusOK, response.Failed{response.Failure, errorMessage})
	// 写法2
	//for _, err2 := range err.(validator.ValidationErrors) {
	//	if err2.Error() != "" {
	//		Exception.ValidBadRequest(c, err2.Translate(global.Translator))
	//		break
	//	}
	//}
}
