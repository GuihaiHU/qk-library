package qutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gogf/gf/frame/g"
)

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// Hash 用Hmac加密
func Hash(key string) string {
	salt := g.Cfg().Get("server.Salt").(string)
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(key))
	return fmt.Sprintf("%02x", h.Sum(nil))
}

func StructToMap(s interface{}) (m g.MapStrAny) {
	j, _ := json.Marshal(&s)
	_ = json.Unmarshal(j, &m)
	return
}

func GetDeepType(typ reflect.Type) reflect.Type {
	resKind := typ.Kind()
	if resKind == reflect.Array || resKind == reflect.Slice || resKind == reflect.Ptr {
		return GetDeepType(typ.Elem())
	} else {
		return typ
	}
}
