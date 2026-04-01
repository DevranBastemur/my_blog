package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

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
	if expiry, exists := lockoutExpiry[ip]; exists {
		if time.Now().Before(expiry) {
			authMutex.Unlock()
			remaining := int(time.Until(expiry).Minutes()) + 1
			renderTemplate(w, "login.page.tmpl", &TemplateData{Error: "Çok fazla hatalı deneme! Sistem " + strconv.Itoa(remaining) + " dakika kilitlendi."})
			return
		} else {
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

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "auth", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func isAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return false
	}
	return cookie.Value == "true"
}
