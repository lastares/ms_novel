package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 自定义一个 Gorm 的数据类型，用于时间存储
type Datetime struct {
	time.Time
}

func (t Datetime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

//  为 JSONTime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t Datetime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 为 JSONTime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Datetime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Datetime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
