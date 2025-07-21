package workerpool

import (
	"log"
	"sync"

	"github.com/scrivx/file-processor/internal/processor"
)

type Result struct {
	FilePath string
	Output   interface{}
	Err      error
}

// StartWorkerPool lanza N workers que procesan archivos en paralelo.
// wg es un puntero a un WaitGroup que será marcado cuando los workers terminen.
func StartWorkerPool(workerCount int, fileChan <-chan string, resultChan chan<- Result,wg *sync.WaitGroup, processorFactory func() (processor.FileProcessor, error),
) {
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int){
			defer wg.Done()
			for filePath := range fileChan {
				fp, err := processorFactory()
				if err != nil{
					log.Printf("[worker %d] Error al crear el procesador: %s", workerID, err)
					resultChan <- Result{FilePath: filePath, Err: err}
					continue
				}
				output, err := fp.Process(filePath)
				resultChan <- Result{FilePath: filePath, Output: output, Err: err}
			}
		} (i)
	}
}
