package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	filescanner "github.com/scrivx/file-processor/internal/file_scanner"
	"github.com/scrivx/file-processor/internal/processor"
	"github.com/scrivx/file-processor/internal/workerpool"
)

func main() {
	// Flags CLI
	dir := flag.String("dir", "", "Directorio con archivos a procesar")
	workers := flag.Int("workers", 4, "Numero de gorutines worker")
	procType := flag.String("type", "wordCount", "Tipo de procesaminto: wordCount, uppercase, checksum")
	flag.Parse()

	if *dir == "" {
		fmt.Println("Uso: --dir <ruta> --workers <N> --type <tipo>")
		os.Exit(1)
	}

	// Crear Canales
	fileChan := make(chan string, 100)
	resultChan := make(chan workerpool.Result, 100)

	// Escaneo del directorio
	go func() {
		err := filescanner.ScanDir(*dir, fileChan)
		if err != nil {
			log.Fatalf("Error escaneando directorio: %v", err)
		}
		close(fileChan) // Cerramos para que los workers terminen cuando terminen de leer
	}()

	// crear el pool de workers
	var wg sync.WaitGroup
	workerpool.StartWorkerPool(*workers, fileChan, resultChan, &wg, func()(processor.FileProcessor, error){
		// Si se requiere checksum con algoritmo, se puede usar algo como:
		if *procType == "checksum" {
			return processor.CkecksumProcessor{Algorithm: "sha256"}, nil
		}
		return processor.NewProcessor(*procType)
	})

	// Esperar a que los workers terminen
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Mostrar resultados
	for result := range resultChan {
		if result.Err != nil{
			log.Printf("❌ Error procesando %s: %v", result.FilePath, result.Err)
		} else {
			fmt.Printf("✅ %s -> %v\n", result.FilePath, result.Output)
		}
	}
}