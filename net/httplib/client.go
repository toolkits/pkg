package httplib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
)

func PostJSON(url string, timeout int, v interface{}) (response []byte, err error) {
	var bs []byte
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	bs, err = json.Marshal(v)
	if err != nil {
		return
	}

	bf := bytes.NewBuffer(bs)

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	var resp *http.Response
	resp, err = client.Post(url, "application/json", bf)
	if err != nil {
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code not equals 200")
	}

	return
}
