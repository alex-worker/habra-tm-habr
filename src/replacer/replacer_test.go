package replacer

import "testing"

func Test_DoSomeTM(t *testing.T) {
	testStr := "Наверное, в любых проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи внешних сервисов и т.д.."
	actual := DoSomeTM(testStr)
	expected := "Наверное, в любых™ проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи™ внешних сервисов и т.д.."
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
