package errors

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type PageError struct {
	Message string
}

func (p PageError) Error() string {
	return p.Message
}

func (p PageError) String() string {
	return p.Message
}

func Bomb(format string, a ...interface{}) {
	panic(PageError{Message: fmt.Sprintf(format, a...)})
}

func Dangerous(v interface{}) {
	if v == nil {
		return
	}

	switch t := v.(type) {
	case string:
		if t != "" {
			panic(PageError{Message: t})
		}
	case error:
		panic(PageError{Message: t.Error()})
	}
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func MaybePanic(err error) {
	_, whichFile, line, _ := runtime.Caller(1)
	arr := strings.Split(whichFile, "/")
	file := arr[len(arr)-1]
	t := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	if err != nil {
		panic(Error{Msg: err.Error(), Time: t, File: file, Line: line})
	}
}

func Panic(msg string) {
	_, whichFile, line, _ := runtime.Caller(1)
	arr := strings.Split(whichFile, "/")
	file := arr[len(arr)-1]
	t := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	panic(Error{Msg: msg, Time: t, File: file, Line: line})
}
