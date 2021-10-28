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

type PICReq struct {
	CardNumber string `json:"cardNumber"`
	Name       string `json:"name"`
	CouponId   string `json:"couponId"`
	Sign       string `json:"sign"`
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
	RefreshToken string `json:"refreshToken"`
	SubaoId      string `json:"subaoId"`
}

// PULReq PostUserLogin - PUL
type PULReq struct {
	PhoneNumber  string   `json:"phoneNumber"`
	PasswordSha1 string   `json:"passwordSha1"`
	SubaoId      string   `json:"subaoId"`
	Code         string   `json:"Code"`
	EntranceType string   `json:"entranceType"`
	CouponList   []string `json:"couponList"`
	TaskList     []string `json:"taskList"`
}
