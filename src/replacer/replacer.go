package replacer

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const RunesInWorld = 5
const StringTemplate = `(\s|^|\pP)[А-Яа-яA-Za-z]{CHARS_NUM}(\pP|\s)`

func DoSomeTM(str string) string {

	compiledStr := strings.Replace(StringTemplate, "CHARS_NUM", strconv.Itoa(RunesInWorld), 1)
	reg := regexp.MustCompile(compiledStr)

	replFn := func(s string) string {

		log.Printf("{%s}", s)

		runes := []rune(s)
		runeCount := len(runes)

		lastRune := runes[runeCount-1]
		runes[runeCount-1] = '™'

		return string(runes) + string(lastRune)
	}

	result := reg.ReplaceAllStringFunc(str, replFn)

	return result
}
