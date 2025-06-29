package images

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ScanAndStoreImages сканирует директорию и добавляет изображения в БД
func ScanAndStoreImages(dir string, dbConn *sql.DB) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Ошибка чтения папки %s: %v", dir, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		switch ext {
		case ".jpg", ".jpeg", ".png", ".webp", ".gif":
			// Сохраняем только имя файла, без директории
			log.Printf("Найдена картинка: %s", name)
			err := InsertImage(dbConn, name)
			if err != nil {
				log.Printf("Ошибка вставки %s: %v", name, err)
			}
		}
	}
}

func InsertImage(dbConn *sql.DB, path string) error {
	query := `INSERT OR IGNORE INTO images (path) VALUES (?)`
	_, err := dbConn.Exec(query, path)
	return err
}
