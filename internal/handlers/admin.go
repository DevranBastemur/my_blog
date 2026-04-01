package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (app *App) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

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

	imagePath, err := uploadImage(r, "image")
	if err != nil {
		log.Println("Görsel yükleme hatası:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err = app.Blogs.Insert(title, content, imagePath)
	if err != nil {
		log.Println("Ekleme hatası:", err)
		http.Error(w, "Kaydedilemedi", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *App) EditPostPage(w http.ResponseWriter, r *http.Request) {
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

	blog, err := app.Blogs.Get(id)
	if err != nil {
		http.Error(w, "Yazı bulunamadı", http.StatusNotFound)
		return
	}

	renderTemplate(w, "edit.page.tmpl", &TemplateData{Blog: blog})
}

func (app *App) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	imagePath, err := uploadImage(r, "image")
	if err != nil {
		log.Println("Görsel yükleme hatası:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	existingImage := r.FormValue("existing_image")

	if imagePath == "" {
		imagePath = existingImage
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}

	err = app.Blogs.Update(id, title, content, imagePath)
	if err != nil {
		log.Println("Güncelleme hatası:", err)
		http.Error(w, "Güncellenemedi", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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

func uploadImage(r *http.Request, formKey string) (string, error) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		return "", err
	}

	file, header, err := r.FormFile(formKey)
	if err != nil {
		if err == http.ErrMissingFile {
			return "", nil
		}
		return "", err
	}
	defer file.Close()

	if header.Size > 5*1024*1024 {
		return "", fmt.Errorf("dosya çok büyük (maksimum 5MB olmalı)")
	}

	buff := make([]byte, 512)
	_, _ = file.Read(buff)
	file.Seek(0, io.SeekStart)

	mimeType := http.DetectContentType(buff)
	var ext string
	switch mimeType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	case "image/webp":
		ext = ".webp"
	default:
		return "", fmt.Errorf("GÜVENLİK İHLALİ: Sadece resim dosyası yükleyebilirsiniz")
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := fmt.Sprintf("./ui/static/uploads/%s", fileName)
	dbPath := fmt.Sprintf("/static/uploads/%s", fileName)

	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return dbPath, nil
}
