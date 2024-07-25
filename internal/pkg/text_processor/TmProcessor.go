package text_processor

import (
	"regexp"
	"strconv"
	"strings"
)

type TmProcessor struct {
	reg *regexp.Regexp
}

func NewTmProcessor(runesInWord int) TextProcessorInterface {
	return &TmProcessor{
		reg: setRunesInWorld(runesInWord),
	}
}

const StringTemplate = `(\s|^|\pP)[А-Яа-яA-Za-z]{:RunesInWorld}(\pP|\s)`

func setRunesInWorld(runesInWord int) *regexp.Regexp {
	compiledStr := strings.Replace(StringTemplate, ":RunesInWorld", strconv.Itoa(runesInWord), 1)
	return regexp.MustCompile(compiledStr)
}

func (p *TmProcessor) ProcessText(str string) string {
	return p.reg.ReplaceAllStringFunc(str, doReplace)
}

func doReplace(s string) string {
	runes := []rune(s)
	runeCount := len(runes)

	lastRune := runes[runeCount-1]
	runes[runeCount-1] = '™'

	return string(runes) + string(lastRune)
}
