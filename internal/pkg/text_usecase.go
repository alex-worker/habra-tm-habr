package pkg

import "habra-tm-habr/internal/pkg/text_processor"

type TextProcessorInterface interface {
	text_processor.TextProcessorInterface
}

func NewTmProcessor(runesInWord int) TextProcessorInterface {
	return text_processor.NewTmProcessor(runesInWord)
}
