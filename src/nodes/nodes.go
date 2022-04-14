package nodes

import (
	"bytes"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"io"
)

//func AddSomeTM(reader *strings.Reader) (string, error) {
//	doc, err := html.Parse(reader)
//	if err != nil {
//		return "", err
//	}
//
//	nodeAddTM(doc, replacer.DoSomeTM)
//
//	var buffer = new(bytes.Buffer)
//	err = html.Render(buffer, doc)
//	if err != nil {
//		return "", err
//	}
//	return string(buffer.Bytes()), nil
//}

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

func NodeAddTM(node *html.Node, replacerFn func(string) string) {
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
			node.Data = replacerFn(node.Data)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		NodeAddTM(child, replacerFn)
	}
}
