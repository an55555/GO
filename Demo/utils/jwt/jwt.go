package jwt

import (
	//"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

// SALT 密钥
const SALT = "secret"

// Header 消息头部
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// PayLoad 负载
type PayLoad struct {
	Sub      string `json:"sub"`
	UserName string `json:"userName"`
	Admin    bool   `json:"admin"`
	// Expire int `json:"exp"`
}

// JWT 完整的本体
type JWT struct {
	Header    `json:"header"`
	PayLoad   `json:"payload"`
	Signature string `json:"signature"`
}

func NewPayload(userName string) PayLoad {
	return PayLoad{
		Sub:      "user",
		UserName: userName,
		Admin:    false,
	}
}

// Encode 将 json 转成符合 JWT 标准的字符串
func (jwt *JWT) Encode() string {
	header, err := json.Marshal(jwt.Header)
	checkError(err)
	headerString := base64.StdEncoding.EncodeToString(header)
	payload, err := json.Marshal(jwt.PayLoad)
	payloadString := base64.StdEncoding.EncodeToString(payload)
	checkError(err)

	format := headerString + "." + payloadString
	signature := getHmacCode(format)

	return format + "." + signature
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(SALT))
	h.Write([]byte(s))
	key := h.Sum(nil)
	return hex.EncodeToString(key)
}

// Decode 验证 jwt 签名是否正确,并将json内容解析出来
func (jwt *JWT) Decode(code string) bool {

	arr := strings.Split(code, ".")
	if len(arr) != 3 {
		return false
	}

	// 验证签名是否正确
	format := arr[0] + "." + arr[1]
	signature := getHmacCode(format)
	if signature != arr[2] {
		return false
	}

	header, err := base64.StdEncoding.DecodeString(arr[0])
	checkError(err)
	payload, err := base64.StdEncoding.DecodeString(arr[1])
	checkError(err)

	json.Unmarshal(header, &jwt.Header)
	json.Unmarshal(payload, &jwt.PayLoad)

	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func JwtDecode(token string) (JWT, bool) {
	jwt := JWT{}
	if jwt.Decode(token) {
		return jwt, false
	}
	return jwt, true
}

func JwtCode(p PayLoad) string {
	jwt := JWT{}
	jwt.Header = Header{"HS256", "JWT"}
	jwt.PayLoad = p
	result := jwt.Encode()
	return result
}
