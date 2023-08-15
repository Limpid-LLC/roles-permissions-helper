package roles_permissions_helper

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// GetLibraryRootPath возвращает путь к корневой директории библиотеки.
func GetLibraryRootPath() (string, error) {
	// Получаем путь к файлу, где находится текущая функция.
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to determine file path")
	}

	// Вычисляем абсолютный путь к директории, содержащей файл с текущей функцией.
	dir := filepath.Dir(file)

	// Ищем корневую директорию библиотеки, поднимаясь по директориям
	// до тех пор, пока не найдем файл с заданным именем (например, go.mod).
	for {
		// Проверяем, существует ли файл go.mod в текущей директории.
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Поднимаемся на уровень выше.
		newDir := filepath.Dir(dir)
		if newDir == dir {
			break // Достигли корневой файловой системы, завершаем цикл.
		}
		dir = newDir
	}

	return "", errors.New("library root not found")
}
