package q

import (
	"strings"
)

type Page struct {
	Current  int `v:"required|integer|min:1#页码不能为空|当前页码必须为整数|当前页码必须为正整数" validate:"required"`
	PageSize int `v:"required|integer|between:1,30#条数不能为空|条数必须为整数|条数必须为 :min 到 :max 之间的正整数" validate:"required"`
}

// 只用于可分页，也可获取全部列表的接口
type PageUnlimited struct {
	// 获取全部列表时传0或不传
	Current int `v:"integer|min:0#当前页码必须为整数|当前页码必须为非负整数"`
	// 获取全部列表时传0或不传
	PageSize int `v:"integer|between:0,30#条数必须为整数|条数必须为 :min 到 :max 之间的整数"`
}

type BatchIds struct {
	// 用逗号分割id
	Ids string `v:"required#ids不能为空" validate:"required" example:"1,2"`
}

func (b BatchIds) GetIdSlice() []string {
	return strings.Split(b.Ids, ",")
}

type IBatchIds interface {
	GetIdSlice() []string
}
