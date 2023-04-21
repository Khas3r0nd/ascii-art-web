package function

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

func CheckHash(check, font string) bool {
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "db448376863a4b9a6639546de113fa6f"
	if check == hashStandard && font == "banners/standard.txt" {
		return true
	}
	if check == hashShadow && font == "banners/shadow.txt" {
		return true
	}
	if check == hashThinkertoy && font == "banners/thinkertoy.txt" {
		return true
	}
	return false
}

func MD5(data string) string {
	bs, err := os.ReadFile(data)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	h := md5.Sum([]byte(string(bs)))
	return fmt.Sprintf("%x", h)
}
