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
	c.HelpInfo = "USAGE: go run main.go domain ${领域}\n领域\t描述\n-----------------\n"
	domain := gcmd.GetArg(2)
	for _, d := range c.Domains {
		d.Register(c)
		if d.DomainName == domain {
			d.Run()
			return
		}
	}
	if domain == "" {
		c.Help()
		return
	}
	panic("不存在domain:" + domain)
}

func (c *DomainCmdService) Help() {
	fmt.Print(c.HelpInfo)
}
