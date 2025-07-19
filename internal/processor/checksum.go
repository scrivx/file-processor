package processor

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ChecksumProcessor calcula el checksum (SHA256 por defecto) de un archivo
type CkecksumProcessor struct{
	 Algorithm string // md5 o sha256
}

func (p CkecksumProcessor) Process(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hashValue string

	switch p.Algorithm {
		case "md5":
			hasher := md5.New()
			if _, err := io.Copy(hasher, file); err != nil {
				return nil, err
			}
			hashValue = hex.EncodeToString(hasher.Sum(nil))
		case "sha256", "":
			hasher := sha256.New()
			if _, err := io.Copy(hasher, file); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("algoritmo de hash no soportado: %s", p.Algorithm)
	}
	return hashValue, nil
}