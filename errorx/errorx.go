package errorx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/i18n"
)

type PageError struct {
	Message string
	Code    int
}

func (p PageError) Error() string {
	return p.Message
}

func (p PageError) String() string {
	return p.Message
}

func Bomb(code int, format string, a ...interface{}) {
	panic(PageError{Code: code, Message: fmt.Sprintf(format, a...)})
}

func BombWithI18n(ctx *gin.Context, code int, format string, a ...interface{}) {
	lang := ctx.GetHeader("X-Language")
	panic(PageError{Code: code, Message: i18n.Sprintf(lang, format, a...)})
}

func Dangerous(v interface{}, code ...int) {
	if v == nil {
		return
	}

	c := 200
	if len(code) > 0 {
		c = code[0]
	}

	switch t := v.(type) {
	case string:
		if t != "" {
			panic(PageError{Code: c, Message: t})
		}
	case error:
		panic(PageError{Code: c, Message: t.Error()})
	}
}

func DangerousWithI18n(ctx *gin.Context, v interface{}, code ...int) {
	if v == nil {
		return
	}

	c := 200
	if len(code) > 0 {
		c = code[0]
	}

	lang := ctx.GetHeader("X-Language")
	switch t := v.(type) {
	case string:
		if t != "" {
			panic(PageError{Code: c, Message: i18n.Sprintf(lang, t)})
		}
	case error:
		panic(PageError{Code: c, Message: i18n.Sprintf(lang, t.Error())})
	}
}
