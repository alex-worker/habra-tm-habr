package replacer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"habra-tm-habr/src/replacer"
	"testing"
)

func TestReplacer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Replacer Suite")
}

var _ = Describe("replacer", Label("DoSomeTM"), func() {
	Describe("DoSomeTM test", func() {
		It("Worked", func() {
			testStr := "Наверное, в любых проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи внешних сервисов и т.д.."
			actual := replacer.DoSomeTM(testStr)
			expected := "Наверное, в любых™ проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи™ внешних сервисов и т.д.."
			Expect(actual).To(Equal(expected))
		})
	})
})
