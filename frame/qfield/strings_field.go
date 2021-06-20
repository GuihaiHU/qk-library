package qfield

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/gogf/gf/util/gconv"
)

type Strings []string

// gorm 自定义结构需要实现 Value Scan 两个方法
// Value 实现方法
func (p Strings) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Strings) Scan(data interface{}) error {
	ids := []byte{}
	gconv.Scan(data, &ids)
	return json.Unmarshal(ids, &p)
}

// 转回原始类型
func (p *Strings) GetSlice() []string {
	return []string(*p)
}
