package processor

import (
	"bufio"
	"os"
)

type WordCountProcessor struct{}

func (wp WordCountProcessor) Process(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan(){
		count++
	}

	if err := scanner.Err(); err !=nil {
		return nil, err
	}

	return count, nil
}