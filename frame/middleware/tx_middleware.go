package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	service "github.com/iWinston/qk-library/frame/qservice"
	"gorm.io/gorm"
)

// 事务
func TX(r *ghttp.Request) {
	db := service.ReqContext.GetDB(r.Context())
	db.Transaction(func(tx *gorm.DB) error {
		service.ReqContext.SetTX(r.Context(), tx)
		r.Middleware.Next()
		return service.ReqContext.GetError(r.Context())
	})
}
