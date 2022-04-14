package replacer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReplacer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Replacer Suite")
}

var _ = Describe("DoSomeTM", func() {
	It("Worked", func() {
		testStr := "Наверное, в любых проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи внешних сервисов и т.д.."
		actual := DoSomeTM(testStr)
		expected := "Наверное, в любых™ проектах есть необходимость использования различных секретных данных - строки подключения к БД, АПИ-ключи™ внешних сервисов и т.д.."
		Expect(actual).To(Equal(expected))
	})
})

var _ = Describe("doReplace", func() {
	It("it must be worked", func() {
		testStr := "Наверное "
		actual := doReplace(testStr)
		expected := "Наверное™ "
		Expect(actual).To(Equal(expected))
	})
	It("it must be worked too", func() {
		testStr := "Наверное"
		actual := doReplace(testStr)
		expected := "Наверно™е"
		Expect(actual).To(Equal(expected))
	})
})
