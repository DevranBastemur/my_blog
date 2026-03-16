package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Home, ana sayfa isteğini karşılar
func Home(w http.ResponseWriter, r *http.Request) {
	// Sadece tam olarak "/" yoluna gelen istekleri kabul et. (Güvenlik için)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Şablon dosyalarımızın yolları
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	// Şablonları parse et
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Şablon hatası:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Template'i çalıştır ve istemciye (w) gönder
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println("Render hatası:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
