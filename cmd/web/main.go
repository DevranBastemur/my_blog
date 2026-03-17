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

	app := &handlers.App{
		Blogs: &models.BlogModel{DB: db},
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Rotalarımız
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.LoginPage(w, r)
		} else if r.Method == http.MethodPost {
			app.LoginPost(w, r)
		}
	})
	mux.HandleFunc("/admin", app.AdminDashboard)
	mux.HandleFunc("/admin/post", app.CreatePost)
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
