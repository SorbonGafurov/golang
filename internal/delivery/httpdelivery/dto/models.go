package dto

// ===== DOMAIN =====
type Request struct {
	Login    string `json:"login" xml:"Login"`
	Password string `json:"password" xml:"Password"`
}

type Response struct {
	ResponseCode    string `json:"responseCode" xml:"ResponseCode"`
	ResponseMessage string `json:"responseMessage" xml:"ResponseMessage"`
}
