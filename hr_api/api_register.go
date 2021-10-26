package hr_api

import (
	"encoding/json"
	"net/http"
)

type RSP map[string]interface{}

func RandCodeRequest(phoneNumber string) RSP {
	url := devServerUrlHead + "/api/v2/android/verification_codes"
	bodyStruct := PVCReq{
		PhoneNumber: phoneNumber,
		UserID:      "",
	}
	bodyJson, _ := json.Marshal(bodyStruct)
	rsp := postAndPrintlnRes(bodyJson, url, http.MethodPost, map[string]string{"Authorization": Authorization})
	return rsp
}

func RegisterWithRandCode(randCode string, phoneNumber string) RSP {
	url := devServerUrlHead + "/api/v2/android/users"
	bodyStruct := PRReq{
		PhoneNumber:      phoneNumber,
		PasswordSha1:     "",
		VerificationCode: randCode,
		SubaoId:          "",
		ReferrerPhone:    "",
	}
	bodyJson, _ := json.Marshal(bodyStruct)
	rsp := postAndPrintlnRes(bodyJson, url, http.MethodPost, map[string]string{"Authorization": Authorization})
	return rsp
}

func RefreshSession(token string, sessionId string, userId string, accessToken string) RSP {
	url := devServerUrlHead + "/api/v2/android/sessions" +
		"?grant_type=refresh_token" + "&sessionId=" + sessionId + "&service=android"
	bodyStruct := PRATReq{
		RefreshToken: token,
		SubaoId: "",
	}
	bodyJson, _ := json.Marshal(bodyStruct)
	rsp := postAndPrintlnRes(bodyJson, url, http.MethodPost, map[string]string{
		"Authorization": Authorization,
		"userId": userId,
		"accessToken": accessToken,
	})
	return rsp
}
