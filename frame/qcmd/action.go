package qcmd

import (
	"github.com/gogf/gf/frame/g"
)

type ActionParam struct {
	Short    string
	Name     string
	Desc     string
	Required bool // 必填
	NoValue  bool // 不需要值模式,例如-f,--force  如果NoValue为true，则忽略Required
}

type Action struct {
	Desc    string
	Params  []ActionParam
	Handler func(param g.MapStrStr) error
}
