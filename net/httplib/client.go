package httplib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
)

func PostJSON(url string, timeout int, v interface{}, headers map[string]string) (response []byte, err error) {
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

	req, err := http.NewRequest("POST", url, bf)
	req.Header.Set("Content-Type", "application/json")
	
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	var resp *http.Response
	resp, err = client.Do(req)
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
