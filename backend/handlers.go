package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandleMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		menuListesi := []Menu{}
		rows, err := db.Query("SELECT id_urun, urun_adi, kategori, fiyat FROM menu")
		if err != nil {
			http.Error(w, "Veritabanı sorgusu başarısız oldu", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var m Menu
			if err := rows.Scan(&m.ID_Urun, &m.UrunAdi, &m.Kategori, &m.Fiyat); err != nil {
				http.Error(w, "Veritabanı taraması başarısız oldu", http.StatusInternalServerError)
				return
			}
			menuListesi = append(menuListesi, m)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Satırları okurken hata oluştu", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(menuListesi)
	}
	if r.Method == http.MethodPost {
		var menu Menu
		err := json.NewDecoder(r.Body).Decode(&menu)
		if menu.UrunAdi == "" || menu.Fiyat <= 0 {
			http.Error(w, "Ürün adı boş olamaz ve fiyat 0'dan büyük olmalıdır", http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
			return
		}
		err = db.QueryRow("SELECT id_urun FROM menu WHERE urun_adi = ?", menu.UrunAdi).Scan(&menu.ID_Urun)
		if err == sql.ErrNoRows {
			_, err = db.Exec("INSERT INTO menu (urun_adi, kategori, fiyat) VALUES (?, ?, ?)", menu.UrunAdi, menu.Kategori, menu.Fiyat)
			if err != nil {
				http.Error(w, "Menü öğesi eklenemedi", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(menu)
			return
		}
		if err != nil {
			http.Error(w, "Menü öğesi zaten mevcut", http.StatusConflict)
			return
		}
	}
}

func HandleMusteri(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		musteriListesi := []Musteri{}
		rows, err := db.Query("SELECT musteri_id, musteri_adi, borc FROM musteriler")
		if err != nil {
			http.Error(w, "Veritabanı sorgusu başarısız oldu", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var m Musteri
			if err := rows.Scan(&m.MusteriID, &m.MusteriAdi, &m.Borc); err != nil {
				http.Error(w, "Veritabanı taraması başarısız oldu", http.StatusInternalServerError)
				return
			}
			musteriListesi = append(musteriListesi, m)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Satırları okurken hata oluştu", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(musteriListesi)
	}

	if r.Method == http.MethodPost {
		var musteri Musteri
		err := json.NewDecoder(r.Body).Decode(&musteri)
		if err != nil {
			http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
			return
		}
		if musteri.MusteriAdi == "" {
			http.Error(w, "Müşteri adı boş olamaz!", http.StatusBadRequest)
			return
		}
		err = db.QueryRow("SELECT musteri_id FROM musteriler WHERE musteri_adi = ?", musteri.MusteriAdi).Scan(&musteri.MusteriID)
		if err == sql.ErrNoRows {
			_, err = db.Exec("INSERT INTO musteriler (musteri_adi, borc) VALUES (?, 0)", musteri.MusteriAdi, musteri.Borc)
			if err != nil {
				http.Error(w, "Müşteri eklenemedi", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(musteri)
			return
		}
		if err == nil {
			http.Error(w, "Müşteri zaten mevcut", http.StatusConflict)
			return
		}
	}
}

func HandleIslem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		islemListesi := []Islem{}
		rows, err := db.Query("SELECT id_islem, musteri_id, urun_adi, tutar, islem_tipi, durum FROM islemler")
		if err != nil {
			http.Error(w, "Veritabanı sorgusu başarısız oldu", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var i Islem
			if err := rows.Scan(&i.ID_Islem, &i.MusteriID, &i.UrunAdi, &i.Tutar, &i.IslemTipi, &i.Durum); err != nil {
				http.Error(w, "Veritabanı taraması başarısız oldu", http.StatusInternalServerError)
				return
			}
			islemListesi = append(islemListesi, i)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Satırları okurken hata oluştu", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(islemListesi)
	}
	if r.Method == http.MethodPost {
		var islem Islem
		err := json.NewDecoder(r.Body).Decode(&islem)
		if err != nil {
			http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
			return
		}
		if islem.MusteriID <= 0 || islem.Tutar <= 0 || islem.IslemTipi != "alisveris" && islem.IslemTipi != "odeme" {
			http.Error(w, "Geçersiz işlem bilgileri", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("INSERT INTO islemler (musteri_id, urun_adi, tutar, islem_tipi) VALUES (?, ?, ?, ?)", islem.MusteriID, islem.UrunAdi, islem.Tutar, islem.IslemTipi)
		if err != nil {
			http.Error(w, "İşlem eklenemedi", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(islem)
	}
}

func HandleHesapDetay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Sadece GET isteği destekleniyor", http.StatusMethodNotAllowed)
		return
	}

	MusteriID := r.URL.Query().Get("musteri_id")
	if MusteriID == "" {
		http.Error(w, "Musteri ID belirtilmedi", http.StatusBadRequest)
		return
	}

	sorgu := `
	SELECT m.musteri_adi, i.urun_adi, i.tutar, i.islem_tipi, i.durum
	FROM islemler i
	JOIN musteriler m ON i.musteri_id = m.id_musteri
	WHERE m.id_musteri = ?`

	rows, err := db.Query(sorgu, MusteriID)
	if err != nil {
		http.Error(w, "Veritabanı sorgusu başarısız oldu", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	hesapDetayListesi := []HesapDetay{}
	for rows.Next() {
		var hd HesapDetay
		if err := rows.Scan(&hd.MusteriAdi, &hd.UrunAdi, &hd.Tutar, &hd.IslemTipi, &hd.Durum); err != nil {
			http.Error(w, "Veritabanı taraması başarısız oldu", http.StatusInternalServerError)
			return
		}
		hesapDetayListesi = append(hesapDetayListesi, hd)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Satırları okurken hata oluştu", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(hesapDetayListesi)
}
