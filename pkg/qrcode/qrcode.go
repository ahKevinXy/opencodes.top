package qrcode

import (
	"bytes"
	"encoding/base64"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/skip2/go-qrcode"
	"unsafe"
)

// string 转base64 图片
func CreateImageToBase64(message string) string {
	png, err := qrcode.Encode(message, qrcode.Medium, 256)
	if err != nil {
		logs.Error("获取文件失败")
		return ""
	}
	buffStore := make([]byte, 5000000)
	base64.StdEncoding.Encode(buffStore, png)
	index := bytes.IndexByte(buffStore, 0)
	baseImage := buffStore[0:index]

	return "data:image/png;base64," + *(*string)(unsafe.Pointer(&baseImage))
}
