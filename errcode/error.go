package errcode

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Code int

func (c Code) Error() string {
	return string(c)
}

// 相对于直接返回code,增加了日志打印
func NewError(code Code, extras ...interface{}) error {
	var originErrString string
	if len(extras) > 0 {
		if err, ok := extras[0].(error); ok {
			originErrString = err.Error()
			extras = extras[1:]
		}
	}
	_, file, line, _ := runtime.Caller(1)
	_, _ = fmt.Fprintf(os.Stderr, "[custom][error] %s %s:%d err[%d] %s\n", time.Now().Format("2006-01-02 15:04:05"), file, line, int(code), originErrString)
	return code
}
