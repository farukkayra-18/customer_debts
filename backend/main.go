package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite" // Yeni pure-Go SQLite sürücümüz
)

func main() {
	// 1. Veritabanı dosyasını aç (dosya yoksa otomatik oluşturur)
	var err error
	db, err = sql.Open("sqlite", "./bufe.db")
	if err != nil {
		log.Fatal("Veritabani hazirlanamadi:", err)
	}
	defer db.Close() // Program kapanırken veritabanı bağlantısını güvenle kapatır

	// 2. Bağlantıyı gerçekten test et
	err = db.Ping()
	if err != nil {
		log.Fatal("Veritabanina bağlanilamadi:", err)
	}

	fmt.Println("Veritabanina basarili bir sekilde bağlanildi ve bufe.db hazir!")

	tablolariOlustur(db)

	http.HandleFunc("/musteri", HandleMusteri)
	http.HandleFunc("/menu", HandleMenu)
	http.HandleFunc("/islem", HandleIslem)
	http.HandleFunc("/hesapdetay", HandleHesapDetay)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
