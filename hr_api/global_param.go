package hr_api

const (
	clientKey        = "sw678c5GoWJakDEnFn"
	imageUrl1        = "D:\\subao\\需求评审\\实名认证2nd\\api_test\\asset\\image\\wangxinping.jpg"
	imageUrl2        = "D:\\subao\\需求评审\\实名认证2nd\\api_test\\asset\\image\\tiny.jpg"
	couponId         = "54bb6c1d7efd36b0a2a47943d5953e1"
	devServerUrlHead = "http://dev.xunyou.mobi"
	wsse             = "UsernameToken Username=\"subaokeji\", " +
		"PasswordDigest=\"uZYYJI62LQhtBf0tnjRdFbC5UlU=\", " +
		"Nonce=\"d36e316282959a9ed4c89851497a717f\", " +
		"Created=\"2003-12-15T14:43:07Z\""
	authorization = "Bearer eyJhbGciOiJSUzUxMiJ9.eyJpYXQiOjE1OTQwOTEwMzMsImV4cCI6MTkwNTEzMTAzMywiY29udGVudCI6In" +
		"tcInRlbmFudFwiOlwiYW5kcm9pZFwiLFwiYWNjb3VudFwiOlwiYzJkNTA0YzYtMWJkZC00MmQyLWFkMTgxXCIsXCJzY29wZXNcIjpbXCJ" +
		"CYWNrYm9uZVwiLFwiUW9zXCIsXCJNdWx0aVwiLFwiSHJcIixcIlZhdWx0XCJdfSJ9.wTHPqXwnv5YIP3-olOQvP1nNFN5XEO7m5o8zRai" +
		"FBKAPjLOjZWL3ugQnOaYPU-IOZCnz-Aq2JG4nhTdfpy42R1ZcR4qqBAsBDq1wT21O8wA426Qi5XxbYhkUxxX_X41xMWx09AlUb0N1qwEn" +
		"CogJRnOrksJOYRaMrwWU5eo6SwY"
)

const (
	VerificationCode = "1234"
)

var (
	FaceComparisonUsers = map[string]faceRecognitionInfo{
		"normal":    {userName: "王新平", idNumber: "330482199508090611", imageUrl: imageUrl1},
		"wrongId":   {userName: "王新平", idNumber: "330482199508090612", imageUrl: imageUrl1},
		"tinyImage": {userName: "王新平", idNumber: "330482199508090612", imageUrl: imageUrl2},
		"illegal":   {userName: "", idNumber: "", imageUrl: ""},
	}
	IDNumberConfirmUsers = map[string]idNumberConfirmInfo{
		"normal":  {userName: "王新平", idNumber: "330482199508090611"},
		"wrongId": {userName: "王新平", idNumber: "330482199508090612"},
		"illegal": {userName: "", idNumber: ""},
	}
)

type faceRecognitionInfo struct {
	userName string
	idNumber string
	imageUrl string
}

type idNumberConfirmInfo struct {
	userName string
	idNumber string
}
