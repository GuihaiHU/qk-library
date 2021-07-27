package q

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/iWinston/qk-library/frame/qservice"
	"github.com/iWinston/qk-library/qutil"

	"github.com/gogf/gf/net/ghttp"
)

func ResponseWithMeta(r *ghttp.Request, err error, data interface{}, total int64) {
	if err != nil {
		qservice.ReqContext.SetError(r.Context(), err)
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok", dataToLowerCamelMap(data), total)
	}
}

func ResponseWithData(r *ghttp.Request, err error, data interface{}) {
	if err != nil {
		qservice.ReqContext.SetError(r.Context(), err)
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok", dataToLowerCamelMap(data))
	}
}

func Response(r *ghttp.Request, err error) {
	if err != nil {
		qservice.ReqContext.SetError(r.Context(), err)
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok")
	}
}

func dataToLowerCamelMap(data interface{}) interface{} {
	typ := reflect.TypeOf(data).Elem()
	if typ != nil && typ.Kind() == reflect.Slice {
		m := qutil.SliceToMap(data)
		lowerCamelMapArr(m)
		return m
	} else {
		m := qutil.StructToMap(data)
		lowerCamelMap(m)
		return m
	}
}

func lowerCamelMap(m map[string]interface{}) {
	for k, v := range m {
		typ := reflect.TypeOf(v)
		if typ != nil && typ.Kind() == reflect.Map {
			lowerCamelMap(v.(map[string]interface{}))
		}
		if typ != nil && typ.Kind() == reflect.Slice {
			lowerCamelMapArr(v.([]interface{}))
		}

		if unicode.IsUpper(rune(k[0])) {
			newKey := strings.ToLower(k[:1]) + k[1:]
			m[newKey] = v
			delete(m, k)
		}
	}
}

func lowerCamelMapArr(ms []interface{}) {
	for _, m := range ms {
		lowerCamelMap(m.(map[string]interface{}))
	}
}

// func scanDataToAddTag(data interface{}, resType reflect.Type) {
// 	resType = qutil.GetDeepType(resType)
// 	for i := 0; i < resType.NumField(); i++ {
// 		item := resType.Field(i)
// 		// 是匿名结构体，递归字段
// 		if item.Anonymous {
// 			scanDataToAddTag(data, item.Type)
// 			continue
// 		}
// 		tag := item.Tag.Get("json")
// 		fmt.Println(tag)
// 		if tag == "" {
// 			item.Tag = `json:"test"`
// 		}
// 	}
// }

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

type ResMeta struct {
	Total    int64 `json:"total"`
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
}

// 数据返回通用JSON数据结构(带分页)
type JsonResponseWithMeta struct {
	Meta ResMeta `json:"meta"`
	JsonResponse
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	jsonResponse := JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	}

	if len(data) > 1 {
		r.Response.WriteJson(JsonResponseWithMeta{
			JsonResponse: jsonResponse,
			Meta: ResMeta{
				Total:    data[1].(int64),
				Current:  r.GetQueryInt64("current"),
				PageSize: r.GetQueryInt64("pageSize"),
			},
		})
	} else {
		r.Response.WriteJson(jsonResponse)
	}
}

func HttpStatus(r *ghttp.Request, err int) {
	if err >= 100 && err < 600 {
		r.Response.WriteHeader(err)
	}
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	HttpStatus(r, err)
	Json(r, err, msg, data...)
	r.Exit()
}
