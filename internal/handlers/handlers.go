package handlers

import (
	"html/template"
	"kisisel-blog/internal/models"
	"log"
	"net/http"
)

type App struct {
	Blogs *models.BlogModel
}

type TemplateData struct {
	Blogs    []*models.BlogPost
	Blog     *models.BlogPost
	Comments []*models.Comment
	Error    string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	files := []string{
		"./ui/html/base.layout.tmpl",
		"./ui/html/" + tmpl,
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Şablon parse hatası:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println("Şablon execute hatası:", err)
	}
}
