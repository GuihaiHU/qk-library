package middleware

import (
	"github.com/iWinston/qk-library/frame/service"
	"gorm.io/gorm"

	"github.com/iWinston/qk-library/frame/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GenRequestContextMiddleware(DB *gorm.DB) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		requestCtx := &model.RequestContext{
			Request: r,
			DB:      DB,
			Data:    make(g.Map),
		}
		service.RequestContext.Init(r, requestCtx)
		r.Middleware.Next()
	}
}
