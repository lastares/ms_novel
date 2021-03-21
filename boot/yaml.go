package boot

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// yaml配置初始化
func yamlInit() {
	// 加载yaml翻译配置文件
	viper.SetConfigName("messages.zh_CN.yaml")
	viper.SetConfigType("yaml")           // REQUIRED if the global file does not have the extension in the name
	viper.AddConfigPath("./translations") // path to look for the global file in
	err := viper.ReadInConfig()           // Find and read the global file
	if err != nil {                       // Handle errors reading the global file
		fmt.Println((err.Error()))
		os.Exit(0)
	}
	fmt.Println("yaml翻译配置初始化完毕...")
}
