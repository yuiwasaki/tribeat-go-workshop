package clog

import (
	"fmt"
	"runtime"
	"time"
)

func Println(a ...interface{}) (n int, err error) {
	now := time.Now()
	_, file, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("file:%s:%d", file, line)
	return fmt.Println(append([]interface{}{now, str}, a...)...)
}
