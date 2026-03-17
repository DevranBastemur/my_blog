package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"kisisel-blog/internal/handlers"
	"kisisel-blog/internal/models"

	_ "modernc.org/sqlite" // SQLite sürücüsü (alt çizgi ile sadece init fonksiyonunu tetikliyoruz)
)

func main() {
	// Veritabanı bağlantısını aç (Eğer dosya yoksa otomatik oluşturur)
	db, err := sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal("Veritabanı açılamadı:", err)
	}
	defer db.Close()

	// Tablomuzu oluştur (Eğer daha önce oluşturulmadıysa)
	createTableSQL := `CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Tablo oluşturulamadı:", err)
	}

	// Uygulama bağımlılıklarını (App struct) kur
	app := &handlers.App{
		Blogs: &models.BlogModel{DB: db},
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handler fonksiyonunu app üzerinden çağırıyoruz
	mux.HandleFunc("/", app.Home)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Sunucu :4000 portunda başlatılıyor...")
	err = srv.ListenAndServe()
	log.Fatal(err)
}
