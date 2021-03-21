package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"ms_novel/response"
	"ms_novel/utils"
	"os"
	"time"
)

func GlobalLoggerMiddleware() gin.HandlerFunc {
	logFileName := os.Getenv("LOG_PATH") + "/" + os.Getenv("APP_ENV") + ".log"
	src, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	loggerTool := logrus.New()
	//设置日志级别
	loggerTool.SetLevel(logrus.DebugLevel)
	//设置输出
	loggerTool.Out = src

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logFileName+".%Y-%m-%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	loggerTool.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		// 把读过的字节流再重新放到body中去，不然后面控制器里接收不到参数了
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		params := string(body)
		defer c.Request.Body.Close()

		//fileName, line, functionName := utils.GetExceptionWhereInfo()
		//处理请求
		responseWrite := &response.BodyLogWriter{Body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = responseWrite
		c.Next()

		// 响应数据处理
		params = utils.TrimSpaceLine(params)
		responseJsonData := responseWrite.Body.String()

		// 响应数据解析到结构体
		var responseData response.ResponseData
		json.Unmarshal(responseWrite.Body.Bytes(), &responseData)

		// 日志格式
		logrusFields := logrus.Fields{
			"url":        c.Request.Host + c.Request.RequestURI,
			"params":     params,
			"method":     c.Request.Method,
			"statusCode": c.Writer.Status(),
			"ip":         c.ClientIP(),
			//"fileName":     fileName,
			//"line":         line,
			//"functionName": functionName,
			// 返回的响应数据和 code,有大用
			"responseCode": responseData.Code,
			"responseData": responseJsonData,
		}

		// 异常只要code不是0和-1日志等级为info就好了，其他的异常等级为error
		if responseData.Code == response.Ok || responseData.Code == response.Failure {
			loggerTool.WithFields(logrusFields).Infoln(responseData.Msg)
		} else {
			loggerTool.WithFields(logrusFields).Errorln(responseData.Msg)
			// todo 这里还可以继续往下拓展，对于生产环境出现的错误异常需要发送邮件、短信提醒、群消息提醒等这些操作来及时提醒开发人员进行处理
		}
	}
}
