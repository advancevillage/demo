package repository

import (
	"business/model"
	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	Customers(offset, limit int) ([]*model.Customer, error)
}

///////////////////////////////////////////////////////////////////////////////////////////
//services
type CustomerService struct {
	Repo CustomerRepository
}

func(s *CustomerService) Customers(offset, limit int) ([]*model.Customer, error) {
	return s.Repo.Customers(offset, limit)
}

///////////////////////////////////////////////////////////////////////////////////////////
//database repository
type CustomerDatabaseRepository struct {
	DB  *gorm.DB
}

func (r *CustomerDatabaseRepository) Customers(offset, limit int) ([]model.Customer, error) {
	return nil, nil
}
///////////////////////////////////////////////////////////////////////////////////////////
//cache repository
type CustomerRedisRepository struct {

}
