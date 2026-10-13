package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	TumMusterileriGetir() ([]Musteri, error)
}

type sqliteRepo struct {
	db *sqlx.DB
}

func (repo *sqliteRepo) TumMusterileriGetir() ([]Musteri, error) {
	var musteriListesi []Musteri
	err := repo.db.Select(
		&musteriListesi,
		"SELECT musteri_id, musteri_adi, borc FROM musteriler",
	)
	if err != nil {
		return nil, err
	}
	return musteriListesi, nil
}

/*
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
*/
