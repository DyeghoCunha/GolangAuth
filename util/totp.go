package util

import (
	"bytes"
	"encoding/base64"
	"image/png"

	//	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTOTPKey(issuer, account string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})
	if err != nil {
		return "", "", err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return "", "", err
	}

	qrBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return qrBase64, key.Secret(), nil
}

func ValidateTOTPCode(secret, code string) bool {
	return totp.Validate(code, secret)
}

func Generate2FASecret(email string) (string, error) {
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "GolangAuth",
		AccountName: email,
	})
	if err != nil {
		return "", err
	}
	return secret.Secret(), nil
}
