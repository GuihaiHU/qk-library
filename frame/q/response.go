package q

import (
	service "github.com/iWinston/qk-library/frame/service"

	"github.com/gogf/gf/net/ghttp"
)

func ResponseWithMeta(r *ghttp.Request, err error, data interface{}, total int64) {
	if err != nil {
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok", data, total)
	}
}

func ResponseWithData(r *ghttp.Request, err error, data interface{}) {
	if err != nil {
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok", data)
	}
}

func Response(r *ghttp.Request, err error) {
	if err != nil {
		service.RequestContext.SetError(r.Context(), err)
		JsonExit(r, 1, err.Error())
	} else {
		JsonExit(r, 0, "ok")
	}
}

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
