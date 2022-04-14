package nodes

import (
	"bytes"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"habra-tm-habr/src/replacer"
	"strings"
)

func AddSomeTM(reader *strings.Reader) (string, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return "", err
	}

	nodeAddTM(doc)

	var buffer = new(bytes.Buffer)
	err = html.Render(buffer, doc)
	if err != nil {
		return "", err
	}
	return string(buffer.Bytes()), nil
}

func nodeAddTM(node *html.Node) {
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
			node.Data = replacer.DoSomeTM(node.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodeAddTM(child)
	}
}
