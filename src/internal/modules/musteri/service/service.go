package service

import "github.com/farukkayra-18/customer_debts/src/internal/modules/musteri/repository"

type Service interface {
	TumMusterileriGetir() ([]repository.Musteri, error)
}

type service struct {
	repo repository.Repository
}

// herhangi bir kontrol yada input olmadığı için bu kısımda service
// boş bulunabilir
func (s *service) TumMusterileriGetir() ([]repository.Musteri, error) {
	return s.repo.TumMusterileriGetir()
}
