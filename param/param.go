package param

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/lwb0214/pkg/errors"
)

func String(r *http.Request, key string, defVal string) string {
	if val, ok := r.URL.Query()[key]; ok {
		if val[0] == "" {
			return defVal
		}
		return strings.TrimSpace(val[0])
	}

	if r.Form == nil {
		errors.Dangerous(r.ParseForm())
	}

	val := r.Form.Get(key)
	if val == "" {
		return defVal
	}

	return strings.TrimSpace(val)
}

func MustString(r *http.Request, key string, displayName ...string) string {
	val := String(r, key, "")
	if val == "" {
		name := key
		if len(displayName) > 0 {
			name = displayName[0]
		}
		errors.Bomb("%s is necessary", name)
	}
	return val
}

func Int64(r *http.Request, key string, defVal int64) int64 {
	raw := String(r, key, "")
	if raw == "" {
		return defVal
	}

	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return defVal
	}

	return val
}

func MustInt64(r *http.Request, key string, displayName ...string) int64 {
	raw := String(r, key, "")
	if raw == "" {
		name := key
		if len(displayName) > 0 {
			name = displayName[0]
		}
		errors.Bomb("%s is necessary", name)
	}

	val, err := strconv.ParseInt(raw, 10, 64)
	errors.Dangerous(err)

	return val
}

func Int(r *http.Request, key string, defVal int) int {
	raw := String(r, key, "")
	if raw == "" {
		return defVal
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		return defVal
	}

	return val
}

func MustInt(r *http.Request, key string, displayName ...string) int {
	name := key
	if len(displayName) > 0 {
		name = displayName[0]
	}

	raw := String(r, key, "")
	if raw == "" {
		errors.Bomb("%s is necessary", name)
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		errors.Bomb("%s should be integer", name)
	}

	return val
}

func Float64(r *http.Request, key string, defVal float64) float64 {
	raw := String(r, key, "")
	if raw == "" {
		return defVal
	}

	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return defVal
	}

	return val
}

func MustFloat64(r *http.Request, key string, displayName ...string) float64 {
	raw := String(r, key, "")
	if raw == "" {
		name := key
		if len(displayName) > 0 {
			name = displayName[0]
		}
		errors.Bomb("%s is necessary", name)
	}

	val, err := strconv.ParseFloat(raw, 64)
	errors.Dangerous(err)

	return val
}

func Bool(r *http.Request, key string, defVal bool) bool {
	raw := String(r, key, "")
	if raw == "true" || raw == "1" || raw == "on" || raw == "checked" || raw == "yes" {
		return true
	} else if raw == "false" || raw == "0" || raw == "off" || raw == "" || raw == "no" {
		return false
	}

	return defVal
}

func MustBool(r *http.Request, key string) bool {
	raw := ""

	if val, ok := r.URL.Query()[key]; ok {
		raw = strings.TrimSpace(val[0])
	} else {
		errors.Bomb("%s is necessary", key)
	}

	if raw == "true" || raw == "1" || raw == "on" || raw == "checked" || raw == "yes" {
		return true
	} else if raw == "false" || raw == "0" || raw == "off" || raw == "" || raw == "no" {
		return false
	} else {
		errors.Bomb("bad request")
	}

	return false
}

func BindJson(r *http.Request, obj interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("Empty request body")
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(body, obj)
	if err != nil {
		return fmt.Errorf("unmarshal body %s err:%v", string(body), err)
	}
	return err
}
