package random

import (
	"bytes"
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
//		  str, err := RandString(5)
//		  if err != nil {
//		   	  panic(err)	// panic or logging
//		  }
//
//	  2. specify the source string
//		  str, err := RandString(10, "abcdef")
//		  if err != nil {
//			  log.Fatal(err)
//		  }
//
func RandString(n int, sources ...string) (string, error) {
	// use default letters
	if len(sources) == 0 {
		sources = append(sources, defaultLetters)
	}

	var letters string
	for _, s := range sources {
		letters += s
	}
	l := len(letters)

	// new rand seed
	rand.Seed(time.Now().UnixNano())

	var bf bytes.Buffer
	var err error
	for ; n > 0; n-- {
		r := letters[rand.Intn(l)]
		if err = bf.WriteByte(r); nil != err {
			break
		}
	}
	return bf.String(), err
}
