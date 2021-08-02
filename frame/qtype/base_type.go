package qtype

import (
	"time"

	"github.com/gogf/gf/util/gconv"
)

func Str(v string) *string {
	return &v
}

func Float32(v interface{}) *float32 {
	t := gconv.Float32(v)
	return &t
}

func Bool(v bool) *bool {
	return &v
}

func Int(v interface{}) *int {
	t := gconv.Int(v)
	return &t
}

func Uint(v interface{}) *uint {
	t := uint(gconv.Int(v))
	return &t
}

func Now() *time.Time {
	t := time.Now()
	return &t
}
