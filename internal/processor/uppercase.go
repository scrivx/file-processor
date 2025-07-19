package processor

import (
	"bufio"
	"os"
	"strings"
)

type UpperCaseProcessor struct{}

func (p UpperCaseProcessor) Process(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return strings.ToUpper(firstLine), nil
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return "", nil // archivo vacio
}