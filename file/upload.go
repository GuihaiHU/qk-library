package file

import (
	"strings"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// SaveLocalFile 保存一个文件到本地
func SaveLocalFile(r *ghttp.Request, dir string) (string, error) {
	return SaveFile(r, "local", dir, "file", "md5")
}

// SaveLocalFiles 保存多个文件到本地
func SaveLocalFiles(r *ghttp.Request, dir string) ([]string, error) {
	return SaveFiles(r, "local", dir, "files", "md5")
}

var renameMap = map[string]func(file *ghttp.UploadFile) (*ghttp.UploadFile, error){
	"md5": renameFileByMd5,
}
var saveMap = map[string]func(file *ghttp.UploadFile, dir string) (string, error){
	"local": saveLocalFile,
}

// SaveFile 保存文件
func SaveFile(r *ghttp.Request, saveDomain, saveDir string, key string, renameMethod string) (string, error) {
	file := r.GetUploadFile(key)
	file, err := renameMap[renameMethod](file)
	if err != nil {
		return "", err
	}
	return saveMap[saveDomain](file, saveDir)
}

// SaveFiles 保存多个文件
func SaveFiles(r *ghttp.Request, saveDomain, saveDir string, key string, renameMethod string) (filePath []string, err error) {
	var p string
	files := r.GetUploadFiles(key)
	for _, file := range files {
		file, err = renameMap[renameMethod](file)
		if err != nil {
			break
		}
		p, err = saveMap[saveDomain](file, saveDir)
		if err != nil {
			break
		}
		filePath = append(filePath, p)
	}
	return
}

// md5重命名
func renameFileByMd5(file *ghttp.UploadFile) (*ghttp.UploadFile, error) {
	infos := strings.Split(file.Filename, ".")
	f, _ := file.Open()
	info := [1024 * 1024 * 10]byte{} // 读取前10M的数据求MD5
	_, err := f.Read(info[:])
	if err != nil {
		return file, err
	}
	file.Filename = gmd5.MustEncryptBytes(info[:])
	if len(infos) >= 2 {
		file.Filename += "." + infos[len(infos)-1]
	}
	return file, err
}

// 保存到本地
// TODO: 慎重调用.go语言并没有提供沙箱机制，很容易获取获取非法路径
func saveLocalFile(file *ghttp.UploadFile, dir string) (string, error) {
	filename, err := file.Save(g.Cfg().GetString("server.ServerRoot") + "/" + dir)
	return "/" + dir + "/" + filename, err
}
