package handlers

import (
	"html/template"
	"kisisel-blog/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type App struct {
	Blogs *models.BlogModel
}

// Şablonlara veri göndermek için kullanacağımız yapı
type TemplateData struct {
	Blogs    []*models.BlogPost
	Blog     *models.BlogPost
	Comments []*models.Comment
	Error    string
}

var (
	failedAttempts = make(map[string]int)
	lockoutExpiry  = make(map[string]time.Time)
	authMutex      sync.Mutex
)

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
		if colon := strings.LastIndex(ip, ":"); colon != -1 {
			ip = ip[:colon]
		}
	}
	return ip
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

func (app *App) LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.page.tmpl", &TemplateData{})
}

func (app *App) LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Geçersiz istek", http.StatusBadRequest)
		return
	}

	ip := getIP(r)

	authMutex.Lock()
	// Kilit kontrolü
	if expiry, exists := lockoutExpiry[ip]; exists {
		if time.Now().Before(expiry) {
			authMutex.Unlock()
			remaining := int(time.Until(expiry).Minutes()) + 1
			renderTemplate(w, "login.page.tmpl", &TemplateData{Error: "Çok fazla hatalı deneme! Sistem " + strconv.Itoa(remaining) + " dakika kilitlendi."})
			return
		} else {
			// Kilit süresi dolmuş, sıfırla
			delete(lockoutExpiry, ip)
			delete(failedAttempts, ip)
		}
	}
	authMutex.Unlock()

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin123" {
		authMutex.Lock()
		delete(failedAttempts, ip)
		delete(lockoutExpiry, ip)
		authMutex.Unlock()

		// Güvenli Cookie oluştur (XSS korumalı)
		cookie := http.Cookie{Name: "auth", Value: "true", Path: "/", HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	authMutex.Lock()
	failedAttempts[ip]++
	attemptsLeft := 5 - failedAttempts[ip]
	errMsg := "Hatalı giriş! Kalan deneme hakkı: " + strconv.Itoa(attemptsLeft)

	if failedAttempts[ip] >= 5 {
		lockoutExpiry[ip] = time.Now().Add(5 * time.Minute)
		errMsg = "Çok fazla hatalı deneme! Sistem 5 dakika kilitlendi."
	}
	authMutex.Unlock()

	renderTemplate(w, "login.page.tmpl", &TemplateData{Error: errMsg})
}

func (app *App) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Admin paneline tüm yazıları yolla
	blogs, err := app.Blogs.All()
	if err != nil {
		log.Println("DB okuma hatası:", err)
	}

	comments, err := app.Blogs.GetAllComments()
	if err != nil {
		log.Println("Yorum okuma hatası:", err)
	}

	renderTemplate(w, "admin.page.tmpl", &TemplateData{Blogs: blogs, Comments: comments})
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

func (app *App) DeletePost(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}

	err = app.Blogs.Delete(id)
	if err != nil {
		log.Println("Silme hatası:", err)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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

	err = app.Blogs.InsertComment(blogID, content)
	if err != nil {
		log.Println("Yorum ekleme hatası:", err)
	}

	http.Redirect(w, r, "/post?id="+blogIDStr, http.StatusSeeOther)
}

func (app *App) DeleteCommentAdmin(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err == nil && id > 0 {
		err = app.Blogs.DeleteComment(id)
		if err != nil {
			log.Println("Yorum silme hatası:", err)
		}
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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
