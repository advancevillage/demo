package repository

import (
	"business/model"
	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	Customers() ([]*model.Customer, error)
}

///////////////////////////////////////////////////////////////////////////////////////////
//services
type CustomerService struct {
	Repo CustomerRepository
}

func(s *CustomerService) Customers() ([]*model.Customer, error) {
	return s.Repo.Customers()
}

///////////////////////////////////////////////////////////////////////////////////////////
//database repository
type CustomerDatabaseRepository struct {
	DB  *gorm.DB
}

func (r *CustomerDatabaseRepository) Customers() ([]model.Customer, error) {
	return nil, nil
}
///////////////////////////////////////////////////////////////////////////////////////////
//cache repository
type CustomerRedisRepository struct {

}
