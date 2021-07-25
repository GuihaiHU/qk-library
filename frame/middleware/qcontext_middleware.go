package middleware

import (
	service "github.com/iWinston/qk-library/frame/qservice"
	"gorm.io/gorm"

	model "github.com/iWinston/qk-library/frame/qmodel"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GenRequestContextMiddleware(DB *gorm.DB) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		requestCtx := &model.QContext{
			Request: r,
			DB:      DB,
			Data:    make(g.Map),
		}
		service.QContext.Init(r, requestCtx)
		r.Middleware.Next()
	}
}
