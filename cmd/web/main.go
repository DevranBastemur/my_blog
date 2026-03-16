package main

import (
	"kisisel-blog/internal/handlers" // Go modül adına göre import (go.mod'da ne yazıyorsa o)
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Statik dosyaları (CSS, JS, Görseller) güvenli bir şekilde sunma
	// http.Dir ile belirlediğimiz klasör dışına çıkılmasını engelliyoruz (Path Traversal koruması)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Rotalarımız
	mux.HandleFunc("/", handlers.Home)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Sunucu :4000 portunda başlatılıyor...")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
