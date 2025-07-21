# 🧰 Parallel File Processor (Go CLI Tool)

Herramienta de línea de comandos escrita en Go para procesar múltiples archivos en paralelo mediante goroutines. Soporta procesamiento de tipo `wordcount`, `uppercase` y `checksum`.

---

## 📝 Uso

### 📦 Wordcount

```bash
go run ./cmd/fileprocessor --dir ./test/testdata --workers 4 --type wordcount
```

### 📦 Uppercase

```bash
go run ./cmd/fileprocessor --dir ./test/testdata --workers 4 --type uppercase
```

### 📦 Checksum

```bash
go run ./cmd/fileprocessor --dir ./test/testdata --workers 4 --type checksum
```

---

## 📝 Ejemplos
