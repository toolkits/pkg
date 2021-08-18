package random

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// default rand source letters
var defaultLetters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandString returns a new string.
// if `source` is empty, the string is automatically generated from `defaultLetters`,
// otherwise its is generated from the `source` string, eg:
//
//	  1. automatically generated from default letters
//		  str, err := RandString(5, "")
//		  if err != nil {
//		   	  panic(err)	// panic or logging
//		  }
//
//	  2. specify the source string
//		  str, err := RandString(10, `abcd`)
//		  if err != nil {
//			  log.Fatal(err)
//		  }
//
func RandString(n int, source string) (str string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	// new rand seed
	rand.Seed(time.Now().UnixNano())

	var bf bytes.Buffer
	if source == "" {
		source = defaultLetters
	}
	l := len(source)
	for ; n > 0; n-- {
		r := source[rand.Intn(l)]
		if err = bf.WriteByte(r); nil != err {
			break
		}
	}
	return bf.String(), nil
}
