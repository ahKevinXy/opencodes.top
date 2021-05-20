package qrcode

import "testing"

func TestCreateImageToBase64(t *testing.T) {
	s := CreateImageToBase64("张三")
	t.Log(s)
}
