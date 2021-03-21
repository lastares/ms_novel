package boot

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// env配置初始化
func envInit() {
	// 全局自动加载env 配置
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Env auto load failed.")
		os.Exit(0)
	}
	fmt.Println("env全局配置初始化完毕...")
}
