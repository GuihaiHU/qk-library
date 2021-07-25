package middleware

import (
	"github.com/iWinston/qk-library/frame/qservice"
	"gorm.io/gorm"

	"github.com/iWinston/qk-library/frame/qmodel"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GenRequestContextMiddleware(DB *gorm.DB) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		requestCtx := &qmodel.QContext{
			Request: r,
			DB:      DB,
			Data:    make(g.Map),
		}
		qservice.QContext.Init(r, requestCtx)
		r.Middleware.Next()
	}
}
