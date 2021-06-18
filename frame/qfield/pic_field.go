package qfield

import "github.com/iWinston/qk-library/qfile"

type Pic string

// Scan 实现方法
func (p *Pic) Scan(value interface{}) (err error) {
	if value != nil {
		str := string(value.([]byte))
		url := qfile.GetImgURL(str)
		*p = Pic(url)
	}
	return
}
