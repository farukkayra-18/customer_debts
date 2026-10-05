package main

type Musteri struct {
	MusteriID  int     `json:"musteri_id"`
	MusteriAdi string  `json:"musteri_adi"`
	Borc       float64 `json:"borc"`
}

type Menu struct {
	ID_Urun  int     `json:"id_urun"`
	UrunAdi  string  `json:"urun_adi"`
	Kategori string  `json:"kategori"`
	Fiyat    float64 `json:"fiyat"`
}

type Islem struct {
	ID_Islem  int     `json:"id_islem"`
	MusteriID int     `json:"musteri_id"`
	UrunAdi   string  `json:"urun_adi"`
	Tutar     float64 `json:"tutar"`
	IslemTipi string  `json:"islem_tipi"`
	Durum     string  `json:"durum"`
}

type HesapDetay struct {
	MusteriAdi string  `json:"musteri_adi"`
	UrunAdi    string  `json:"urun_adi"`
	Tutar      float64 `json:"tutar"`
	IslemTipi  string  `json:"islem_tipi"`
	Durum      string  `json:"durum"`
}
