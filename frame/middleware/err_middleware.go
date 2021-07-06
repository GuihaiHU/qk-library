package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iWinston/qk-library/frame/q"
	"github.com/iWinston/qk-library/notify"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// ErrorHandler 统一错误处理
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 业务panic
		errInfo := strings.Split(err.Error(), "##")
		if len(errInfo) == 2 {
			var errNum int
			errNum, err2 := strconv.Atoi(errInfo[0])
			if err2 == nil {
				r.Response.ClearBuffer()
				r.Response.WriteHeader(200)
				q.JsonExit(r, errNum, errInfo[1])
			}
		}

		// 普通panic
		errDetail := fmt.Sprintf("%+v", err)
		go notify.Notify(g.Cfg().GetString("server.NodeName") + ":\n" + errDetail)
		g.Log("exception").Error(errDetail)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}
