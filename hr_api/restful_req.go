package hr_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postAndPrintlnRes(jsonBody []byte, url string, method string, heads map[string]string) map[string]interface{} {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	for key, value := range heads {
		req.Header.Add(key, value)
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(rsp.Status, string(bs))
	var result map[string]interface{}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil
	}
	return result
}
