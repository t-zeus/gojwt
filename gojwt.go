// author: t-zeus
package gojwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

func GenMac(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func CheckMac(message, messageMac, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMac, expectedMAC)
}

func base64EncodeURL(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

func base64DecodeURL(text string) string {
	ret, _ := base64.URLEncoding.DecodeString(text)
	return string(ret)
}

func GenJWT(v interface{}, key string) string {
	var (
		header    string
		payload   string
		signature string
	)
	hJson, _ := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	header = base64EncodeURL(string(hJson))

	pJson, _ := json.Marshal(v)
	payload = base64EncodeURL(string(pJson))

	signature = SignJWT(header, payload, key)
	return header + "." + payload + "." + base64EncodeURL(signature)
}

func SignJWT(header, payload, key string) string {
	message := []byte(header + "." + payload)
	return string(GenMac(message, []byte(key)))
}

func VerifyJWT(jwt, key string) bool {
	arr := strings.Split(jwt, ".")
	if len(arr) != 3 {
		return false
	}
	message := arr[0] + "." + arr[1]
	return CheckMac([]byte(message), []byte(base64DecodeURL(arr[2])), []byte(key))
}
