package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Gelen HTTP isteklerini yönlendirecek router (mux) nesnemiz
	mux := http.NewServeMux()

	// Şimdilik ana sayfa için test amaçlı basit bir yanıt döndürüyoruz
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Güvenli Kisisel Blog Sistemine Hos Geldiniz!"))
	})

	// Güvenlik odaklı özel sunucu yapılandırması
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Sunucu :4000 portunda güvenli bir şekilde başlatılıyor...")

	// Sunucuyu ayağa kaldırıyoruz, hata olursa uygulamayı durdurup log basıyoruz
	err := srv.ListenAndServe()
	log.Fatal(err)
}
