package str

import (
	"fmt"
	"strconv"
	"strings"
)

func IdsInt64(ids string, sep ...string) []int64 {
	if ids == "" {
		return []int64{}
	}

	s := ","
	if len(sep) > 0 {
		s = sep[0]
	}

	var arr []string

	if s == " " {
		arr = strings.Fields(ids)
	} else {
		arr = strings.Split(ids, s)
	}

	count := len(arr)
	ret := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		if arr[i] != "" {
			id, err := strconv.ParseInt(arr[i], 10, 64)
			if err == nil {
				ret = append(ret, id)
			}
		}
	}

	return ret
}

func IdsString(ids []int64, sep ...string) string {
	count := len(ids)
	arr := make([]string, count)
	for i := 0; i < count; i++ {
		arr[i] = fmt.Sprint(ids[i])
	}

	if len(sep) > 0 {
		return strings.Join(arr, sep[0])
	}

	return strings.Join(arr, ",")
}
