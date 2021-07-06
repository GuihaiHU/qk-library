package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// LogRequest 访问日志
func LogRequest(r *ghttp.Request) {
	start := time.Now()
	r.Middleware.Next()
	duration := time.Since(start).Milliseconds()
	g.Log().Info(r.Method, r.RequestURI, r.Response.Status, fmt.Sprintf("%dms", duration))
	// user := shared.Context.GetUser(r.Context())
	// if user != nil {
	// 	g.Log().Info("user:" + user.Username)
	// }
	info := strings.Replace(r.GetRawString(), "\n", "", -1)
	info = strings.Replace(info, " ", "", -1)
	info = strings.Replace(info, "\t", "", -1)
	g.Log().Info("requestBody:" + info)
	g.Log().Info("response:" + strings.Replace(r.Response.BufferString(), "\n", "", -1))
	g.Log().Info("======================================")
}
