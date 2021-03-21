package model

type BaseModel struct {
}

var (
	COMMON_ONE   = 1  // true
	COMMON_ZERO  = 0  // false
	COMMON_MINUS = -1 // 负数
)

func GetCommonValues() (commonValues []int) {
	commonValues = []int{COMMON_ONE, COMMON_ZERO}
	return
}
