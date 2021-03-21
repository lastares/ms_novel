package boot

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ms_novel/global"
	"os"
	"time"
)

func mysqlConnect() {
	var err error
	// 读取数据库相关配置
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)

	global.Gorm, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Db connection failed, %s", err.Error())
		os.Exit(0)
	}
	// sql 调试模式
	global.Gorm.LogMode(true)

	// 设置table不是负数形式
	global.Gorm.SingularTable(true)

	// 数据库连接池
	global.Gorm.DB().SetMaxIdleConns(10)           // 设置空闲连接池中的最大连接数
	global.Gorm.DB().SetMaxOpenConns(100)          // 设置数据库连接最大打开数
	global.Gorm.DB().SetConnMaxLifetime(time.Hour) // 设置可重用连接的最长时间
	// 获取通用数据库接口
	global.Gorm.DB().Ping()

	fmt.Println("mysql 服务启动成功...")

}

// 关闭数据库
func MysqlClose() {
	global.Gorm.Close()
}
