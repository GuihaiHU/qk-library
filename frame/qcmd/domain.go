package qcmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
)

var red = color.New(color.FgRed).PrintlnFunc()

var green = color.New(color.FgGreen).PrintlnFunc()
var blue = color.New(color.FgBlue).PrintlnFunc()

type Domain struct {
	Description string
	DomainName  string
	Actions     map[string]Action
}

// Help
func (d *Domain) Help() {
	blue(fmt.Sprintf("\n%s帮助文档:", d.Description))
	for name, action := range d.Actions {
		fmt.Printf("%s ", name)
		for _, p := range action.Params {
			fmt.Printf(" -%s %s ", p.Short, p.Name)
		}
		fmt.Printf("\t%s\n", action.Desc)
	}
}

// parseParam 必填 parser *gcmd.Parser, key string
func (s *Domain) parseParam(parser *gcmd.Parser, action Action) g.MapStrStr {
	params := g.MapStrStr{}
	for i := range action.Params {
		key := action.Params[i].Name
		if action.Params[i].NoValue {
			if parser.ContainsOpt(key) {
				params[key] = "true"
			} else {
				params[key] = "false"
			}
		} else {
			value := parser.GetOpt(key)
			if action.Params[i].Required && value == "" {
				panic("缺少参数[" + key + "]")
			} else {
				params[key] = value
			}
		}
	}
	return params
}

// Register
func (d *Domain) Register(c *DomainCmdService) {
	c.HelpInfo += c.HelpInfo + fmt.Sprintf("%s\t%s", d.DomainName, d.Description)
}

// GenHandle 产生帮助函数
func (d *Domain) Run() {
	if action, ok := d.Actions[gcmd.GetArg(3)]; ok {
		parserConfig := g.MapStrBool{}
		for i := range action.Params {
			key := fmt.Sprintf("%s,%s", action.Params[i].Short, action.Params[i].Name)
			parserConfig[key] = !action.Params[i].NoValue
		}
		parser, err := gcmd.Parse(parserConfig)
		if err != nil {
			return
		}

		params := d.parseParam(parser, action)
		err = action.Handler(params)
		if err != nil {
			red(gcmd.GetArgAll(), gcmd.GetOptAll(), err.Error())
			d.Help()
		} else {
			green("操作成功")
		}
	} else {
		d.Help()
	}
}
