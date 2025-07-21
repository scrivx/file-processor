package workerpool

import (
	"fmt"
	"sync"
	"time"

	"github.com/scrivx/file-processor/internal/processor"
)

type Result struct {
	FilePath string
	Output   interface{}
	Err      error
	WorkerID int
}

// StartWorkerPool lanza N workers que procesan archivos en paralelo.
// wg es un puntero a un WaitGroup que será marcado cuando los workers terminen.
func StartWorkerPool(
	workerCount int,
	fileChan <-chan string,
	resultChan chan<- Result,
	wg *sync.WaitGroup,
	processorFactory func() (processor.FileProcessor, error),
) {
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for filePath := range fileChan {
				fmt.Printf("\033[36m[Worker %d] ⏳ Procesando %s\033[0m\n", workerID, filePath)

				start := time.Now()

				fp, err := processorFactory()
				if err != nil {
					resultChan <- Result{FilePath: filePath, Err: err, WorkerID: workerID}
					continue
				}

				output, err := fp.Process(filePath)

				elapsed := time.Since(start)
				if err != nil {
					fmt.Printf("\033[31m[Worker %d] ❌ Error en %s (%s)\033[0m\n", workerID, filePath, err)
				} else {
					fmt.Printf("\033[32m[Worker %d] ✅ Completado %s en %s\033[0m\n", workerID, filePath, elapsed)
				}

				resultChan <- Result{
					FilePath: filePath,
					Output:   output,
					Err:      err,
					WorkerID: workerID,
				}
			}
		}(i)
	}
}