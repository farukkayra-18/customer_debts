package main

// burası model değil bunlar sadece dto, data transfer object
// model olabilmesi için bu verilerin db üzerinde etkisi olması
// lazım. örnek:
//
//	type Musteri struct {
//		MusteriID  int     `json:"musteri_id" db:"musteri_id`
//		MusteriAdi string  `json:"musteri_adi" db:"musteri_id`
//		Borc       float64 `json:"borc" db:musteri_id`
//	}
//
//	hem model hem de dto olarak iş görür ama genel olarak clean
//	kullanımıysa dtolar farklı modeller farklı olmasıdır. Örnek:
//
//	type MusteriCreateRequest struct {
//		MusteriAdi string  `json:"musteri_adi"`
//	}
//	Neden? yeni müşteri oluşturunca borcu temelde sıfırdır, çünkü
//	yeni oluştu. id si uuid.New() seklinde gelecektir, handler içerisine
//	sadece müşteri ismi alması yeterlidir, devamı service katmanında
//	default valuelar olarak verilip sonra da db kısımında yazma silme yapılır
//	bu sefer de elimizde
//
//	type Musteri struct {
//		MusteriID  int     `db:"musteri_id`
//		MusteriAdi string  `db:"musteri_id`
//		Borc       float64 `db:musteri_id`
//	}
//
//	bu sekilde bir model oluşturulur ve artık bu sayede direkt:
//
//	var musteri Musteri
//
//	err := db.Get(
//		&musteri,
//		"SELECT * FROM musteriler WHERE id = $1",
//	42,
//	)
//
//	bu sayede sürekli işlem yükünden kurtulup db içerisindeki verileri
//	hızlı şekilde struct içerisine okuma ve yazma yapabiliriz sql de
//	field isimleriyle tag içerisinde db:"xxx" isimleri aynı olanlar direkt
//	yazılmış olur

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
