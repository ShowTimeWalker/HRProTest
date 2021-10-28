package hr_api

import (
	. "../common_func"
	"encoding/base64"
	"encoding/json"
	"os"
)

func IDNumberConfirm(users map[string]idNumberConfirmInfo, label string, userId string, accessToken string) RSP {
	userName := users[label].userName
	userCardNumber := users[label].idNumber

	md5Sign := "cardNumber=" + userCardNumber +
		"&" + "couponId=" + couponId +
		"&" + "name=" + userName +
		"&" + "key=" + clientKey
	md5Sign = MD5(md5Sign)

	url := devServerUrlHead + "/api/v1/android/users/" + userId + "/identity/confirmation"

	body := PICReq{
		CardNumber: userCardNumber,
		Name:       "王新平",
		CouponId:   couponId,
		Sign:       md5Sign,
	}
	jsonBody, _ := json.Marshal(body)
	rsp := postAndPrintlnRes(jsonBody, url, map[string]string{
		"Authorization": authorization,
		"accessToken":   accessToken,
	})
	return rsp
}

func FaceComparison(users map[string]faceRecognitionInfo, label string, userId string, accessToken string) RSP {
	userName := users[label].userName
	userCardNumber := users[label].idNumber
	imageUrl := users[label].imageUrl

	md5Sign := "cardNumber=" + userCardNumber +
		"&" + "couponId=" + couponId +
		"&" + "name=" + userName +
		"&" + "key=" + clientKey
	md5Sign = MD5(md5Sign)

	imageFile, _ := os.Open(imageUrl)
	defer func(imageFile *os.File) {
		err := imageFile.Close()
		if err != nil {
		}
	}(imageFile)
	imageBuffer := make([]byte, 1000000)
	n, _ := imageFile.Read(imageBuffer)
	imageBase64Str := base64.StdEncoding.EncodeToString(imageBuffer[:n])

	url := devServerUrlHead + "/api/v1/android/users/" + userId + "/identity/face_recognition"
	body := PFRReq{
		CardNumber: userCardNumber,
		Name:       "王新平",
		CouponId:   couponId,
		Sign:       md5Sign,
		Base64Str:  imageBase64Str,
	}
	jsonBody, _ := json.Marshal(body)
	rsp := postAndPrintlnRes(jsonBody, url, map[string]string{
		"Authorization": authorization,
		"accessToken":   accessToken,
	})
	return rsp
}

func GetAccelGamePermission(userId string, accessToken string) RSP {
	url := devServerUrlHead + "/api/v1/android/users/" + userId + "/accel_game/permission"
	rsp := GetAndPrintRes(url, map[string]string{
		"Authorization": authorization,
		"accessToken":   accessToken,
	})
	return rsp
}
