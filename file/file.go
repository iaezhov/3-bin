package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const storageDir = "file"

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(getFilePath(name))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) (bool, error) {
	file, err := os.Create(getFilePath(name))
	if err != nil {
		return false, fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return false, fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return true, nil
}

func IsJSONFile(name string) bool {
	return filepath.Ext(name) == ".json"
}

func getFilePath(name string) string {
	return filepath.Join(storageDir, name)
}
