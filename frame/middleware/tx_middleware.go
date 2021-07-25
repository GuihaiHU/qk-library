package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	service "github.com/iWinston/qk-library/frame/qservice"
	"gorm.io/gorm"
)

// 事务
func TX(r *ghttp.Request) {
	db := service.QContext.GetDB(r.Context())
	db.Transaction(func(tx *gorm.DB) error {
		service.QContext.SetTX(r.Context(), tx)
		r.Middleware.Next()
		return service.QContext.GetError(r.Context())
	})
}
