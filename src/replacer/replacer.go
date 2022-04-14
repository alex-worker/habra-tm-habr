package replacer

import (
	"regexp"
	"strconv"
	"strings"
)

const RunesInWorld = 5
const StringTemplate = `(\s|^|\pP)[А-Яа-яA-Za-z]{CHARS_NUM}(\pP|\s)`

func doReplace(s string) string {
	runes := []rune(s)
	runeCount := len(runes)

	lastRune := runes[runeCount-1]
	runes[runeCount-1] = '™'

	return string(runes) + string(lastRune)
}

func DoSomeTM(str string) string {

	compiledStr := strings.Replace(StringTemplate, "CHARS_NUM", strconv.Itoa(RunesInWorld), 1)
	reg := regexp.MustCompile(compiledStr)

	result := reg.ReplaceAllStringFunc(str, doReplace)

	return result
}
