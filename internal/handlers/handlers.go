package handlers

import (
	"html/template"
	"kisisel-blog/internal/models"
	"log"
	"net/http"
)

// App struct'ı, handler'larımızın ihtiyaç duyduğu bağımlılıkları (veritabanı gibi) tutar
type App struct {
	Blogs *models.BlogModel
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Şimdilik sadece veritabanı bağlantımızı test ediyoruz, verileri sonraki adımda HTML'e basacağız
	_, err := app.Blogs.Latest()
	if err != nil {
		log.Println("Veritabanından yazılar çekilemedi:", err)
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Şablon hatası:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println("Render hatası:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
