package main

import (
	. "./common_func"
	. "./hr_api"
)

func main() {
	phoneNumber := GenRandPhone("194")
	// 手机号码验证注册并登录
	RandCodeRequest(phoneNumber)
	registerRsp := RegisterWithRandCode(VerificationCode, phoneNumber)

	// 刷新登录
	refreshToken := registerRsp["sessionInfo"].(map[string]interface{})["refreshToken"].(string)
	sessionId := registerRsp["sessionInfo"].(map[string]interface{})["sessionId"].(string)
	userId := registerRsp["userInfo"].(map[string]interface{})["userId"].(string)
	accessToken := registerRsp["sessionInfo"].(map[string]interface{})["accessToken"].(string)
	refreshRsp := RefreshSession(refreshToken, sessionId, userId, accessToken)

	// 人像对比验证
	accessToken = refreshRsp["accessToken"].(string)
	FaceComparison(Users, "tinyImage", userId, accessToken)
}
