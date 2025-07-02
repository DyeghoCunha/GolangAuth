package controller

type Setup2FAResponse struct {
	QRCode string `json:"qr_code"`
	Secret string `json:"secret"`
}
