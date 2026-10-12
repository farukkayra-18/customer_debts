package repository

type Musteri struct {
	MusteriID  int     `db:"musteri_id"`
	MusteriAdi string  `db:"musteri_adi"`
	Borc       float64 `db:"borc"`
}
