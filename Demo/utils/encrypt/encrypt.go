package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

const SALT = "@#@#"

func EncryptSailt(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	md5String := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, SALT)
	io.WriteString(h, md5String)
	md5Result := fmt.Sprintf("%x", h.Sum(nil))
	return md5Result
}
