package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Yeni pure-Go SQLite sürücümüz
)

var db *sql.DB

func tablolariOlustur(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS musteriler (
    	musteri_id INTEGER PRIMARY KEY AUTOINCREMENT,
    	musteri_adi TEXT NOT NULL,
    	borc REAL NOT NULL
	);`)

	if err != nil {
		log.Fatal("Musteriler tablosu oluşturulamadı:", err)
	}

	fmt.Println("Musteriler tablosu basarili bir sekilde olusturuldu veya zaten mevcut.")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS menu (
    id_urun INTEGER PRIMARY KEY AUTOINCREMENT,
    urun_adi TEXT NOT NULL,
    kategori TEXT NOT NULL,
    fiyat REAL NOT NULL
	);`)
	if err != nil {
		log.Fatal("Menü tablosu oluşturulamadı:", err)
	}
	fmt.Println("Menü tablosu hazır.")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS islemler (
    	id_islem INTEGER PRIMARY KEY AUTOINCREMENT,
    	musteri_id INTEGER NOT NULL,
    	urun_adi TEXT,
    	tutar REAL NOT NULL,
    	islem_tipi TEXT NOT NULL,
    	durum TEXT DEFAULT 'aktif'
		);`)
	if err != nil {
		log.Fatal("İşlemler tablosu oluşturulamadı:", err)
	}
	fmt.Println("İşlemler tablosu hazır.")

}
