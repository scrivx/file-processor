package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/scrivx/file-processor/internal/processor"
	"github.com/stretchr/testify/assert"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmp := t.TempDir()
	path := filepath.Join(tmp, "tempfile.txt")
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("no se pudo crear archivo temporal: %v", err)
	}
	return path
}

func TestWordCountProcessor(t *testing.T) {
	content := "hola mundo\nesto es una prueba"
	file := createTempFile(t, content)
	defer os.Remove(file)

	wp := processor.WordCountProcessor{}
	result, err := wp.Process(file)
	assert.NoError(t, err)
	assert.Equal(t, 6, result)
}