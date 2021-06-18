package field

import "codeup.aliyun.com/sevenfifteen/sevenfifteenBoilerplate/go-library/file"

type Pic string

// Scan 实现方法
func (p *Pic) Scan(value interface{}) (err error) {
	if value != nil {
		str := string(value.([]byte))
		url := file.GetImgURL(str)
		*p = Pic(url)
	}
	return
}
