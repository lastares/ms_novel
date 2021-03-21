package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

// 获取异常信息所在文件
func GetExceptionWhereInfo() (fileName string, line int, functionName string) {
	pc, fileName, line, ok := runtime.Caller(3)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}
	return fileName, line, functionName
}
