package model

// ===== DOMAIN =====
type Request struct {
	Login    string `json:"login" xml:"Login" binding:"required"`
	Password string `json:"password" xml:"Password" binding:"required"`
}

type Response struct {
	ResponseCode    string `json:"responseCode" xml:"ResponseCode"`
	ResponseMessage string `json:"responseMessage" xml:"ResponseMessage"`
}

type RequestTest struct {
	Login    string `json:"login" xml:"Login"`
	Password string `json:"password" xml:"Password"`
}

type ResponseTest struct {
	ResponseCode    string `json:"responseCode" xml:"ResponseCode"`
	ResponseMessage string `json:"responseMessage" xml:"ResponseMessage"`
}
