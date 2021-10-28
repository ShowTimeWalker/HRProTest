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
	userId := registerRsp["userInfo"].(map[string]interface{})["userId"].(string)
	refreshToken := registerRsp["sessionInfo"].(map[string]interface{})["refreshToken"].(string)
	sessionId := registerRsp["sessionInfo"].(map[string]interface{})["sessionId"].(string)
	accessToken := registerRsp["sessionInfo"].(map[string]interface{})["accessToken"].(string)
	refreshRsp := RefreshSession(refreshToken, sessionId, userId, accessToken)

	// 用户登出
	sessionId = refreshRsp["sessionId"].(string)
	accessToken = refreshRsp["accessToken"].(string)
	Logout(userId, accessToken, sessionId)

	// 一键登录
	loginRSP := Login(1, "", VerificationCode, phoneNumber)

	// 姓名身份证号认证
	userId = loginRSP["userInfo"].(map[string]interface{})["userId"].(string)
	accessToken = loginRSP["sessionInfo"].(map[string]interface{})["accessToken"].(string)
	IDNumberConfirm(IDNumberConfirmUsers, "normal", userId, accessToken)
	// 查询用户权限
	GetAccelGamePermission(userId, accessToken)

	// 人像对比验证
	FaceComparison(FaceComparisonUsers, "normal", userId, accessToken)
	// 查询用户权限
	GetAccelGamePermission(userId, accessToken)
}
