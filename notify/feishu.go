package notify

import (
	"github.com/gogf/gf/frame/g"
)

// Notify 通知飞书
func Notify(info string) {
	g.Client().Get(g.Cfg().GetString("feishu.ErrNotify"), g.Map{
		"msg_type": "text",
		"content": g.Map{
			"text": info,
		},
	})
}
