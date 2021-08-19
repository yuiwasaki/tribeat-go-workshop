package clog

import (
	"fmt"
	"runtime"
	"time"
)

func Println(a ...interface{}) (n int, err error) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	_, file, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("file:%s:%d", file, line)
	return fmt.Println(append([]interface{}{now, str}, a...)...)
}

func Hoge() bool {
	return true
}
