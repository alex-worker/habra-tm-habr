package parser

import (
	"bytes"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"io"
)

func BytesToHTML(reader io.Reader) (*html.Node, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func HTMLToBytes(doc *html.Node) ([]byte, error) {
	var buffer = new(bytes.Buffer)
	err := html.Render(buffer, doc)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Update(node *html.Node, updateFn func(string) string) {
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
			node.Data = updateFn(node.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		Update(child, updateFn)
	}
}
