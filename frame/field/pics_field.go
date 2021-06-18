package field

import (
	"database/sql/driver"
	"encoding/json"

	"codeup.aliyun.com/sevenfifteen/sevenfifteenBoilerplate/go-library/file"

	"github.com/gogf/gf/util/gconv"
)

type Pics []string

// gorm 自定义结构需要实现 Value Scan 两个方法
// Value 实现方法
func (p Pics) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Pics) Scan(data interface{}) error {
	pics := []byte{}
	gconv.Scan(data, &pics)
	err := json.Unmarshal(pics, &p)
	if err != nil {
		return err
	}
	if p != nil {
		pArr := *p
		for i := 0; i < len(pArr); i++ {
			pArr[i] = file.GetImgURL(pArr[i])
		}
	}
	return nil
}

// 转回原始类型
func (p *Pics) GetSlice() []string {
	return []string(*p)
}
