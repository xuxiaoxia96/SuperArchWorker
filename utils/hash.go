package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func Sha256(input []byte) string {
	h := sha256.New()
	h.Write(input)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256FromStr(input string) string {
	h := sha256.New()
	h.Write([]byte(input))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func ImgMatchHash(imgStr, hashid string) bool{
	var b64Str string
	tmp := strings.Split(imgStr, ",")
	if len(tmp) == 1 {
		b64Str = tmp[0]
	}else if len(tmp) > 1 {
		b64Str = tmp[1]
	}else {
		return false
	}

	data, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return false
	}
	fmt.Printf(Sha256(data))
	if Sha256(data) == hashid {
		return true
	}else {
		return false
	}
}
