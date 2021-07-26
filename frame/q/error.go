package q

import (
	"fmt"
	"regexp"

	"github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/errors/gerror"
)

func OptimizeDbErr(err error) error {
	if err, ok := err.(*mysql.MySQLError); ok {
		switch err.Number {
		case 1062:
			m := regexp.MustCompile(`Duplicate entry '(.*)' for.*`).FindStringSubmatch(err.Message)
			return gerror.New(fmt.Sprintf("%s 已存在", m[1]))
		case 1054:
			panic(err)
		}
	}
	return err
}
