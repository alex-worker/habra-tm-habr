package replacer

import (
	"testing"
)

func TestDoSomeTM(t *testing.T) {
	testStr := "Наверное, в любых проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи внешних сервисов и т.д.."
	actual := DoSomeTM(testStr)
	expected := "Наверное, в любых™ проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи™ внешних сервисов и т.д.."
	if actual != expected {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
}
