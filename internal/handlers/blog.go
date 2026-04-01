package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	blogs, err := app.Blogs.Latest()
	if err != nil {
		log.Println("DB Hatası:", err)
		http.Error(w, "Sunucu Hatası", http.StatusInternalServerError)
		return
	}

	data := &TemplateData{Blogs: blogs}
	renderTemplate(w, "home.page.tmpl", data)
}

func (app *App) ViewPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	blog, err := app.Blogs.Get(id)
	if err != nil {
		http.Error(w, "Yazı bulunamadı", http.StatusNotFound)
		return
	}

	renderTemplate(w, "post.page.tmpl", &TemplateData{Blog: blog})
}

func (app *App) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metod izin verilmiyor", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	blogIDStr := r.FormValue("blog_id")
	content := r.FormValue("content")

	blogID, err := strconv.Atoi(blogIDStr)
	if err != nil || content == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	if len([]rune(content)) > 250 {
		http.Error(w, "Yorum 250 karakterden uzun olamaz", http.StatusBadRequest)
		return
	}

	err = app.Blogs.InsertComment(blogID, content)
	if err != nil {
		log.Println("Yorum ekleme hatası:", err)
	}

	http.Redirect(w, r, "/post?id="+blogIDStr, http.StatusSeeOther)
}
