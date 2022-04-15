package nodes

import (
	"strings"
	"testing"
)

func Test_BytesToHTML_HTMLToBytes(t *testing.T) {
	fixture := `<html><head><title>Hello</title></head><body><h1>Привет!</h1></body></html>`

	myHTML, err := BytesToHTML(strings.NewReader(fixture))
	if err != nil {
		t.Error(err)
		return
	}

	myBytes, err := HTMLToBytes(myHTML)
	if err != nil {
		t.Error(err)
		return
	}

	actual := string(myBytes)

	if actual != fixture {
		t.Errorf("Actual: %s expected: %s", actual, fixture)
	}
}
