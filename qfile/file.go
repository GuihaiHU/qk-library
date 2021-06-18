package qfile

import "github.com/gogf/gf/frame/g"

// GetImgURL 获取图片的url
func GetImgURL(path string) string {
	return g.Cfg().GetString("server.AssetUrl") + path
}

// TransImg 图片转换
func TransImg(entity g.Map, src, dest string) g.Map {
	entity[dest] = GetImgURL(entity[src].(string))
	return entity
}

// GetLocalFileAbsPath 获取本地文件的绝对路径
func GetLocalFileAbsPath(path string) string {
	return g.Cfg().GetString("server.ServerRoot") + path
}
