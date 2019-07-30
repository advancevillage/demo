package repository

import (
	"business/model"
	"github.com/jinzhu/gorm"
	"util"
)
const (
	CustomerTable = "customers"
)

type CustomerRepository interface {
	Customers(offset, limit int) ([]*model.Customer, int, error)
}

///////////////////////////////////////////////////////////////////////////////////////////
//services
type CustomerService struct {
	Repo CustomerRepository
}

func(s *CustomerService) Customers(offset, limit int) ([]*model.Customer, int, error) {
	return s.Repo.Customers(offset, limit)
}

///////////////////////////////////////////////////////////////////////////////////////////
//database repository
type CustomerDatabaseRepository struct {
	DB  *gorm.DB
}

func (r *CustomerDatabaseRepository) Customers(offset, limit int) (customers []*model.Customer, total int, err error) {
	err = r.DB.Table(CustomerTable).Count(&total).Error
	if err != nil {
		return
	}
	min := util.MinInt(total, limit)
	customers = make([]*model.Customer, 0, min)
	err = r.DB.Table(CustomerTable).Find(&customers).Limit(limit).Offset(offset * limit).Error
	return
}
///////////////////////////////////////////////////////////////////////////////////////////
//cache repository
type CustomerRedisRepository struct {

}
