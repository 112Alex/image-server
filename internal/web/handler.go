package web

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"image-server/internal/db"
	"net/http"
)

// ServeIndex отдает главную HTML-страницу
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("D:/GOLANG/image-server/internal/web/templates/index.html")
	if err != nil {
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
