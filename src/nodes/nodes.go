package nodes

import (
	"bytes"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"habra-tm-habr/src/replacer"
	"strings"
)

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
			node.Data = replacer.DoSomeTM(node.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodeReplacer(child)
	}
}
