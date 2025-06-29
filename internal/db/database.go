package db

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite" // Используем sqlite-драйвер (совместим с stdlib)
)

// Image структура представляет запись изображения в БД
type Image struct {
	ID        int       `json:"id"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

// InitDB открывает/создаёт базу и применяет миграции
func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Ошибка открытия базы: %v", err)
	}

	// Создание таблицы, если не существует
	schema := `
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Ошибка создания схемы БД: %v", err)
	}

	return db
}

// InsertImage добавляет путь изображения, если его ещё нет в базе
func InsertImage(db *sql.DB, imgPath string) error {
	query := `INSERT OR IGNORE INTO images (path, created_at) VALUES (?, ?)`
	_, err := db.Exec(query, imgPath, time.Now())
	return err
}

// GetAllImages возвращает список изображений в порядке добавления
func GetAllImages(db *sql.DB) ([]Image, error) {
	rows, err := db.Query(`SELECT id, path, created_at FROM images ORDER BY created_at ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var img Image
		err := rows.Scan(&img.ID, &img.Path, &img.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}
