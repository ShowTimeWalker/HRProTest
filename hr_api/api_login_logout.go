package hr_api

import "encoding/json"

func Logout(userId string, accessToken string, sessionId string) RSP {
	url := devServerUrlHead + "/api/v2/android/sessions/" + sessionId + "?service="
	header := map[string]string{
		"userId":        userId,
		"accessToken":   accessToken,
		"Authorization": authorization,
		"X-WSSE":        wsse,
	}
	rsp := DeleteAndPrintRes(url, header)
	return rsp
}

func Login(method int, password string, randCode string, phoneNumber string) RSP {
	url := devServerUrlHead + "/api/v2/android/sessions"
	if method == 1 {
		url = url + "?grant_type=one_click_login" + "&channel_type=" + "&entrance=" + "&service="
	} else if method == 2 {
		url = url + "?grant_type=short_message" + "&channel_type=" + "&entrance=" + "&service="
	} else {
		url = url + "?grant_type=password" + "&entrance=" + "&service="
	}
	body := PULReq{
		PhoneNumber:  phoneNumber,
		PasswordSha1: password,
		SubaoId:      "",
		Code:         randCode,
		EntranceType: "",
		CouponList:   []string{},
		TaskList:     []string{},
	}
	jsonBody, _ := json.Marshal(body)
	rsp := postAndPrintlnRes(jsonBody, url, map[string]string{
		"Authorization": authorization,
		"X-WSSE":        wsse,
	})
	return rsp
}
