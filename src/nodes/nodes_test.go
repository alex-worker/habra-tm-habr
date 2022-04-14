package nodes

import (
	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/html"
	"regexp"
	"strings"
	"testing"
)

func TestReplacer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nodes Suite")
}

func Test_nodeReplace(t *testing.T) {
	tests := []struct {
		name        string
		mockFunc    func()
		query       string
		expected    string
		mustBeEqual bool
	}{
		{
			name:        "must be worked on p tag",
			mockFunc:    func() {},
			query:       "<p>Приве т!</p>",
			expected:    "<p>Приве™ т!</p>",
			mustBeEqual: true,
		},
		{
			name:        "must be worked on span tag",
			mockFunc:    func() {},
			query:       "<span>Приве т!</span>",
			expected:    "<span>Приве™ т!</span>",
			mustBeEqual: true,
		},
		{
			name:        "must be not worked on unknown tag",
			mockFunc:    func() {},
			query:       "<tag>Приве т!</tag>",
			expected:    "<tag>Приве™ т!</tag>",
			mustBeEqual: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			query := strToHtml(tc.query)
			nodeAddTM(query)
			actual := htmlToStr(query)

			var result bool
			if tc.mustBeEqual {
				result = actual == tc.expected
			} else {
				result = actual != tc.expected
			}
			if !result {
				tt.Errorf("Error expected {%v} but actual {%v}", tc.expected, actual)
			}
		})
	}
}

func htmlToStr(node *html.Node) string {
	var buffer = new(bytes.Buffer)
	err := html.Render(buffer, node)
	if err != nil {
		panic(err)
	}
	respStr := string(buffer.Bytes())
	reg := regexp.MustCompile(`<body>.+</body>`)
	response := reg.FindString(respStr)
	response = strings.ReplaceAll(response, `<body>`, ``)
	response = strings.ReplaceAll(response, `</body>`, ``)
	return response
}

func strToHtml(s string) *html.Node {
	template := `<html><head><title>My title</title></head><body>%s</body></html>`
	input := fmt.Sprintf(template, s)
	result, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	return result
}
