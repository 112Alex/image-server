package web

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"image-server/internal/db"
	"log"
	"net/http"
)

// ServeIndex отдает главную HTML-страницу
// Важно: сервер должен запускаться из корня проекта!
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	const relTmplPath = "internal/web/templates/index.html"
	tmpl, err := template.ParseFiles(relTmplPath)
	if err != nil {
		log.Printf("Ошибка шаблона (%s): %v", relTmplPath, err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// APIHandlerImages возвращает список изображений в формате JSON
func APIHandlerImages(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images, err := db.GetAllImages(dbConn)
		if err != nil {
			http.Error(w, "Ошибка получения изображений", http.StatusInternalServerError)
			return
		}

		// Возвращаем как JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
	}
}
