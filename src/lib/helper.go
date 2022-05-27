package lib

import (
	"regexp"
	"strings"
)

func MakeSlug(text string) string {
	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}
