package processor

import "fmt"

func NewProcessor(processorType string) (FileProcessor, error) {
	switch processorType {
	case "wordcount":
		return WordCountProcessor{}, nil
	case "uppercase":
		return UpperCaseProcessor{}, nil
	case "ckecksum":
		return CkecksumProcessor{}, nil
	default:
		return nil, fmt.Errorf("tipo de procesador no soportado: %s", processorType)
	}
}