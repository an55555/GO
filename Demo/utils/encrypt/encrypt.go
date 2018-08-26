package encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

const SALT = "@#@#"

// salt-md5方案
func EncryptSailt(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	md5String := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, SALT)
	io.WriteString(h, md5String)
	md5Result := fmt.Sprintf("%x", h.Sum(nil))
	return md5Result
}

// 专家方案
func Encrypt(s string) string {
	dk, _ := scrypt.Key([]byte(s), []byte(SALT), 16384, 8, 1, 32)
	return base64.StdEncoding.EncodeToString(dk)
}
