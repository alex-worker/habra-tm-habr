package replacer

import "testing"

func Test_DoSomeTM(t *testing.T) {
	testStr := "В своей книге автор рассказывает о жизни в Венесуэле, в которой огромные запасы нефти сочетаются с нищетой населения и полным отсутствием перспектив, и о причинах экономического коллапса этой страны."
	actual := DoSomeTM(testStr)
	expected := "В своей книге автор рассказывает о жизни в Венесуэле, в которой огромные запасы™ нефти сочетаются с нищетой населения и полным™ отсутствием перспектив, и о причинах экономического коллапса этой страны™."
	if actual != expected {
		t.Errorf("Actual %v but expected %v", actual, expected)
	}
}

func Test_doReplace(t *testing.T) {
	tests := []struct {
		inputStr string
		expected string
	}{
		{
			inputStr: "Наверное ",
			expected: "Наверное™ ",
		},
		{
			inputStr: "Наверное",
			expected: "Наверно™е",
		},
	}
	for _, test := range tests {
		actual := doReplace(test.inputStr)
		if actual != test.expected {
			t.Errorf("Actual %v but expected %v", actual, test.expected)
		}
	}
}
