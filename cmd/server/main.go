package main

import (
	"fmt"
	"image-server/internal/db"
	"image-server/internal/images"
	"image-server/internal/web"
	"log"
	"net/http"
	"os"
)

func ensureDirExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Не удалось создать папку %s: %v", path, err)
		}
	}
}

func main() {
	const dbPath = "data/images.db"
	const staticDir = "static"

	ensureDirExists("data")
	dbConn := db.InitDB(dbPath)
	ensureDirExists(staticDir)
	images.ScanAndStoreImages(staticDir, dbConn)

	// Сканирование изображений в static/
	images.ScanAndStoreImages(staticDir, dbConn)

	// Роуты
	http.HandleFunc("/", web.ServeIndex)
	http.HandleFunc("/api/images", web.APIHandlerImages(dbConn))

	// Отдача статических файлов
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Сервер запущен: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
