package middleware

import (
	"github.com/iWinston/qk-library/frame/service"

	"github.com/gogf/gf/net/ghttp"
	"gorm.io/gorm"
)

// 事务
func TX(r *ghttp.Request) {
	db := service.RequestContext.GetDB(r.Context())
	db.Transaction(func(tx *gorm.DB) error {
		service.RequestContext.SetTX(r.Context(), tx)
		r.Middleware.Next()
		return service.RequestContext.GetError(r.Context())
	})
}
