package qcmd

import (
	"fmt"

	"github.com/gogf/gf/os/gcmd"
)

type DomainCmdService struct {
	Domains  []Domain
	HelpInfo string
}

// Run 发布命令
func (c *DomainCmdService) Run() {
	domain := gcmd.GetArg(2)
	for _, d := range c.Domains {
		d.Register(c)
		if d.DomainName == domain {
			d.Run()
			return
		}
	}
	panic("不存在domain:" + domain)
}

func (c *DomainCmdService) Help() {
	fmt.Print("go run main.go domain 参数")
	fmt.Print(c.HelpInfo)
}
