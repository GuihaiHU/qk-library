package qfield

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/iWinston/qk-library/qfile"
)

type Pics []string

// gorm 自定义结构需要实现 Value Scan 两个方法
// Value 实现方法
func (p Pics) Value() (driver.Value, error) {
	if p == nil {
		p = []string{}
	}
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Pics) Scan(data interface{}) error {
	err := json.Unmarshal(data.([]byte), &p)
	if err != nil {
		return err
	}
	if p != nil {
		pArr := *p
		for i := 0; i < len(pArr); i++ {
			pArr[i] = qfile.GetImgURL(pArr[i])
		}
	}
	return nil
}

// 转回原始类型
func (p *Pics) GetSlice() []string {
	return []string(*p)
}
