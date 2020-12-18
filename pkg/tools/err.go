package tools

import (
	"fmt"
	"runtime"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckErr 错误处理
func CheckErr(err error) {
	checkErr(err)
}

func trace(msg string, args ...interface{}) string {
	msg = fmt.Sprintf(msg, args...)
	logs := []string{msg, ""}
	n := 1
	for {
		n++
		pc, file, line, ok := runtime.Caller(n)
		if !ok {
			break
		}
		f := runtime.FuncForPC(pc)
		name := f.Name()
		if strings.HasPrefix(name, "runtime.") {
			continue
		}
		logs = append(logs, fmt.Sprintf("(%s:%d) %s", file, line, name))
	}

	return strings.Join(logs, "\n")
}

// Trace 查找是那行出的错
func Trace(msg string, args ...interface{}) string {
	return trace(msg, args)
}
