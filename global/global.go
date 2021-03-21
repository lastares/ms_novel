package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	Gorm       *gorm.DB
	Gin        *gin.Engine
	Validate   *validator.Validate
	LoggerTool *logrus.Logger
	Translator ut.Translator
)
