package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	secret = "hello"
)

//密码加密
func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256(append([]byte(str), []byte(secret)...)))
}

//将时间戳转为时间
func SwitchTimeStampToStr(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
