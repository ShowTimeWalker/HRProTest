package hr_api

// PVCReq PostVerificationCode - PVC
type PVCReq struct {
	PhoneNumber string `json:"phoneNumber"`
	UserID      string `json:"userID"`
}

// PRReq PostRegisterReq - PR
type PRReq struct {
	PhoneNumber      string `json:"phoneNumber"`
	PasswordSha1     string `json:"passwordSha1"`
	VerificationCode string `json:"verificationCode"`
	SubaoId          string `json:"subaoId"`
	ReferrerPhone    string `json:"referrerPhone"`
}

// PFRReq PostFaceRecognitionConfirm - PFR
type PFRReq struct {
	CardNumber string `json:"cardNumber"`
	Name       string `json:"name"`
	CouponId   string `json:"couponId"`
	Sign       string `json:"sign"`
	Base64Str  string `json:"base64Str"`
}

// PRATReq PostRefreshAccessToken - PRAT
type PRATReq struct {
	RefreshToken string `json:"refreashToken"`
	SubaoId      string `json:"subaoId"`
}
