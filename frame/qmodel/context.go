package qmodel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gorm.io/gorm"
)

const (
	// 上下文变量存储键名，前后端系统共享
	ContextKey = "ReqContext"
)

// 请求上下文结构
// ghttp.Request.Context() 是一个宽泛的容器，而ReqContext是包含了常见的Request,User等数据的精确容器，用于api,service, dao各层之间的数据传输
// ReqContext 常见用法是在中间件中把它初始化它，然后把它放入ghttp.Request.Context(),需要用的时候通过shared.Context精准拿出需要的数据
type ReqContext struct {
	Request *ghttp.Request // 当前Request对象
	DB      *gorm.DB       // 未开启事务的DB
	TX      *gorm.DB       // 开启事务的DB
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
	Err     error          // 保存错误
}
