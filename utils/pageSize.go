package utils

import (
	"os"
	"strconv"
)

// 获取分页数据
func GetPageData(page, pageSize int) (int, int, int) {
	var zero int
	if page == zero {
		page, _ = strconv.Atoi(os.Getenv("PAGE"))
	}

	if pageSize == zero {
		pageSize, _ = strconv.Atoi(os.Getenv("PAGESIZE"))
	}

	offset := (page - 1) * pageSize
	return page, pageSize, offset
}
