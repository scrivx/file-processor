package processor

type FileProcessor interface {
	Process(filePath string) (interface{}, error)
}