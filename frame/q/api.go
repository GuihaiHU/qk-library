package q

import (
	"reflect"
	"strings"

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
			panic(err.Error())
		}
	}

	dtoType := reflect.TypeOf(param).Elem()
	dtoValue := reflect.ValueOf(param).Elem()

	for i := 0; i < dtoType.NumField(); i++ {
		itemType := dtoType.Field(i)
		itemValue := dtoValue.Field(i)

		// 通过ctx标签获取
		ctxTag := itemType.Tag.Get("ctx")
		if ctxTag != "" {
			arr := strings.Split(ctxTag, ".")
			ctxTagName := arr[0]
			ctxVar := r.GetCtxVar(ctxTagName).Interface()
			ctxVarRef := reflect.ValueOf(ctxVar)
			if len(arr) == 1 {
				itemValue.Set(ctxVarRef)
			} else {
				ctxFieldName := arr[1]
				itemValue.Set(ctxVarRef.Elem().FieldByName(ctxFieldName))
			}
		}
	}
}
