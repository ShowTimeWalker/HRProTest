package hr_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postAndPrintlnRes(body []byte, url string, heads map[string]string) RSP {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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
	var result RSP
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil
	}
	return result
}

func GetAndPrintRes(url string, heads map[string]string) RSP {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
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
	var result RSP
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil
	}
	return result
}

func DeleteAndPrintRes(url string, heads map[string]string) RSP {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
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
	var result RSP
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil
	}
	return result
}
