package nodes

import (
	"bytes"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const RunesInWorld = 5
const StringTemplate = `(\s|^|\pP)[А-Яа-яA-Za-z]{CHARS_NUM}(\pP|\s)`

func AddSomeTM(bytesArr []byte) (string, error) {
	strBody := string(bytesArr)

	doc, err := html.Parse(strings.NewReader(strBody))
	if err != nil {
		return "", err
	}

	return addSomeTM(doc)
}

func addSomeTM(root *html.Node) (string, error) {
	nodeReplacer(root)
	var buffer = new(bytes.Buffer)
	err := html.Render(buffer, root)
	if err != nil {
		return "", err
	}
	return string(buffer.Bytes()), nil
}

func nodeReplacer(node *html.Node) {
	if node.Type == html.TextNode {
		parentDataName := node.Parent.Data

		proceedNode := []string{
			"p",
			"span",
			"h1",
			"a",
			"div",
			"em",
		}

		nodeIsInArray := slices.Contains(proceedNode, parentDataName)
		if nodeIsInArray {
			node.Data = doSomeTM(node.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodeReplacer(child)
	}
}

func doSomeTM(str string) string {

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
