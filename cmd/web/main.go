package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"kisisel-blog/internal/handlers"
	"kisisel-blog/internal/models"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createCommentsSQL := `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		blog_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE
	);`
	_, err = db.Exec(createCommentsSQL)
	if err != nil {
		log.Fatal(err)
	}

	app := &handlers.App{
		Blogs: &models.BlogModel{DB: db},
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Rotalarımız
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/post", app.ViewPost)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.LoginPage(w, r)
		} else if r.Method == http.MethodPost {
			app.LoginPost(w, r)
		}
	})
	mux.HandleFunc("/comment", app.AddComment)
	mux.HandleFunc("/admin", app.AdminDashboard)
	mux.HandleFunc("/admin/post", app.CreatePost)
	mux.HandleFunc("/admin/delete", app.DeletePost)
	mux.HandleFunc("/admin/delete-comment", app.DeleteCommentAdmin)
	mux.HandleFunc("/logout", app.Logout)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Sunucu :4000 portunda başlatılıyor...")
	log.Fatal(srv.ListenAndServe())
}
