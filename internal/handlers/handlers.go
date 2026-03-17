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

// Şablonlara veri göndermek için kullanacağımız yapı
type TemplateData struct {
	Blogs []*models.BlogPost
	Error string
}

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

func (app *App) LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.page.tmpl", nil)
}

func (app *App) LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Geçersiz istek", http.StatusBadRequest)
		return
	}

	// Basit bir güvenlik önlemi (Hardcoded parola, eğitim/proje için uygundur)
	if r.FormValue("password") == "admin123" {
		// Güvenli Cookie oluştur (XSS korumalı)
		cookie := http.Cookie{Name: "auth", Value: "true", Path: "/", HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "login.page.tmpl", &TemplateData{Error: "Hatalı parola!"})
}

func (app *App) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "admin.page.tmpl", nil)
}

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err := app.Blogs.Insert(title, content)
	if err != nil {
		log.Println("Ekleme hatası:", err)
		http.Error(w, "Kaydedilemedi", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "auth", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Yardımcı Fonksiyonlar
func isAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return false
	}
	return cookie.Value == "true"
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	files := []string{
		"./ui/html/" + tmpl,
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Şablon parse hatası:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println("Şablon execute hatası:", err)
	}
}
