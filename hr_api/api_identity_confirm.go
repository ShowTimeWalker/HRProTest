package hr_api

import (
	. "../common_func"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

func FaceComparison(users map[string]faceRecognitionInfo, label string, userId string, accessToken string) RSP{
	userName := users[label].userName
	userCardNumber := users[label].idNumber
	imageUrl := users[label].imageUrl

	md5Sign := "cardNumber=" + userCardNumber +
		"&" + "couponId=" + CouponId +
		"&" + "name=" + userName +
		"&" + "key=" + ClientKey
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

	PFRUrl := "http://dev.xunyou.mobi/api/v1/android/users/" + userId + "/identity/face_recognition"
	PFRBody := PFRReq{
		CardNumber: userCardNumber,
		Name:       "王新平",
		CouponId:   CouponId,
		Sign:       md5Sign,
		Base64Str:  imageBase64Str,
	}
	PFRJsonBody, _ := json.Marshal(PFRBody)
	rsp := postAndPrintlnRes(PFRJsonBody, PFRUrl, http.MethodPost, map[string]string{
		"Authorization": Authorization,
		"accessToken":   accessToken,
	})
	return rsp
}
