package repository

import (
	"business/model"
	"github.com/jinzhu/gorm"
	"util"
)
const (
	CustomerTable = "customers"
)

///////////////////////////////////////////////////////////////////////////////////////////
//database repository
type CustomerDatabaseRepository struct {
	DB  *gorm.DB
}

func (r *CustomerDatabaseRepository) Customers(offset, limit int) (customers []*model.Customer, total int, err error) {
	err = r.DB.Table(CustomerTable).Count(&total).Error
	if err != nil || total == 0 {
		return nil, 0, nil
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
