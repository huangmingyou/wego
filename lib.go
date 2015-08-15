package wxapi

// @基本的功能函数

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Encode() {
	fmt.Println("encodeing!")
}
