package q

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetIdFormReq(r *ghttp.Request) (id uint) {
	if id = r.GetUint("id"); id == 0 {
		JsonExit(r, 1, "未获得id参数")
	}
	return
}

func AssignParamFormReq(r *ghttp.Request, param interface{}) {
	// 先从入参中获取
	if err := r.Parse(param); err != nil {
		if reflect.TypeOf(err).Elem().Name() == "validationError" {
			JsonExit(r, 400, err.Error())
		} else {
			g.Log("exception").Error(fmt.Sprintf("%+v", err))
			JsonExit(r, 500, err.Error())
		}
	}

	// 解析标签
	parseParamByTag(r, param)
}

func parseParamByTag(r *ghttp.Request, param interface{}) {
	dtoType := reflect.TypeOf(param).Elem()
	dtoValue := reflect.ValueOf(param).Elem()

	for i := 0; i < dtoType.NumField(); i++ {
		itemType := dtoType.Field(i)
		itemValue := dtoValue.Field(i)

		// 忽略 swagger:ignore 的值
		ignoreTag := itemType.Tag.Get("swaggerignore")
		if ignoreTag == "true" && !itemValue.IsNil() {
			err := fmt.Errorf("不允许传递%s参数", itemType.Name)
			g.Log("exception").Error(err)
			JsonExit(r, 400, err.Error())
		}

		// 设置默认值
		defaultTag, isTagExisted := itemType.Tag.Lookup("default")
		if isTagExisted {
			defaultValue := getDefaultValue(r, itemType.Type, defaultTag)
			itemValue.Set(reflect.ValueOf(defaultValue))
		}

		// 通过ctx标签获取
		ctxTag := itemType.Tag.Get("ctx")
		if ctxTag != "" {
			arr := strings.Split(ctxTag, ".")
			ctxTagName := arr[0]
			ctxVar := r.GetCtxVar(ctxTagName).Interface()
			ctxVarRef := reflect.ValueOf(ctxVar)
			if ctxVarRef.IsNil() {
				err := fmt.Errorf("获取不到%s的值", ctxTag)
				g.Log("exception").Error(err)
				JsonExit(r, 500, err.Error())
			}
			if len(arr) == 1 {
				itemValue.Set(ctxVarRef)
			} else {
				ctxFieldName := arr[1]
				itemValue.Set(ctxVarRef.Elem().FieldByName(ctxFieldName))
			}
		}
	}
}

func getDefaultValue(r *ghttp.Request, typ reflect.Type, defaultTag string) (defaultValue interface{}) {
	switch typ.Kind() {
	case reflect.String:
		defaultValue = defaultTag
	case reflect.Bool:
		if value, err := strconv.ParseBool(defaultTag); err != nil {
			g.Log("exception").Error(err)
			JsonExit(r, 500, err.Error())
		} else {
			defaultValue = value
		}
	case reflect.Uint:
		if value, err := strconv.ParseUint(defaultTag, 10, 64); err != nil {
			g.Log("exception").Error(err)
			JsonExit(r, 500, err.Error())
		} else {
			defaultValue = value
		}
	case reflect.Float32:
		if value, err := strconv.ParseFloat(defaultTag, 32); err != nil {
			g.Log("exception").Error(err)
			JsonExit(r, 500, err.Error())
		} else {
			defaultValue = value
		}
	case reflect.Float64:
		if value, err := strconv.ParseFloat(defaultTag, 64); err != nil {
			g.Log("exception").Error(err)
			JsonExit(r, 500, err.Error())
		} else {
			defaultValue = value
		}
	default:
		err := errors.New("不支持的默认值类型" + typ.String())
		g.Log("exception").Error(err)
		JsonExit(r, 500, err.Error())
	}
	return
}
