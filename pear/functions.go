package pear

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// It will panic if such config value don't exists
func GetStr(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	return value
}

// It will panic if such config value don't exists
// or the value converts to integer unsuccessfully
func GetInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	output, err := strconv.Atoi(value)
	if err != nil {
		log.Panic(`Environmental variable [` + key + `] is not an integer`)
	}

	return output
}

// It will panic if such config value don't exists
func GetBytes(key string) []byte {
	value := os.Getenv(key)
	if value == "" {
		log.Panic(`Environmental variable [` + key + `] don't exists`)
	}
	return []byte(value)
}
func PasswordMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func InputCheck(src string) string {
	//Remove SCRIPT
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//Remove STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	src = strings.TrimSpace(src)
	return src
}
func InputCheckUrl(data url.Values) map[string]string {
	paramMap := make(map[string]string, 0)
	for k, v := range data {
		if len(v) == 1 && len(v[0]) != 0 {
			paramMap[k] = v[0]
			//Remove SCRIPT
			re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
			paramMap[k] = re.ReplaceAllString(paramMap[k], "")
			//Remove STYLE
			re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
			paramMap[k] = re.ReplaceAllString(paramMap[k], "")
			paramMap[k] = strings.TrimSpace(paramMap[k])
		} else {
			break
		}
	}
	return paramMap
}
