// Author By ChaiHanLin. Date 10/14/19.
// Contact me <ChaiHanLin@blued.com>
package util

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

func RecordPanicStack(e interface{}, n string) {
	TrueOrFalse(CreateNewFile(fmt.Sprintf("./runtime/_panic_%s", n), []byte(fmt.Sprintf("%v\n", e)+string(debug.Stack()))))
}

func TrueOrFalse(b interface{}) bool {
	if b != nil {
		fn, file, line, _ := runtime.Caller(1)
		Stderr("%s:%s:%d %v ", file, runtime.FuncForPC(fn).Name(), line, b)
		return false
	}
	return true
}

func NoError(e error) {
	if e != nil {
		panic(e)
	}
}

func Stdout(format string, v ...interface{}) {
	if _, err := fmt.Fprintf(os.Stdout, fmt.Sprintf("[SRV] %s %s \n", time.Now().Format(time.RFC3339), format), v...); err != nil {
		panic(err)
	}
}

func Stderr(format string, v ...interface{}) {
	if _, err := fmt.Fprintf(os.Stderr, fmt.Sprintf("[ERR] %s %s \n", time.Now().Format(time.RFC3339), format), v...); err != nil {
		panic(err)
	}
}

//EOF
