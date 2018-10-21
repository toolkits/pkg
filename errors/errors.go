package errors

import "fmt"

type PageError struct {
	Message string
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