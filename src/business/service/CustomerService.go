//@brief: 数据库接口层 定义底层持久层的指向
//@note:  应该有且依赖business/model, business/repository
package service

import (
	"business/model"
	"business/repository"
	"pool"
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

// @brief: CustomerService读取数据库的类型
// @param: b boolean
// @eg:
//   b = true （write)  获取写连接
//   b = false (read)   获取读连接
func NewCustomerService(b bool) *CustomerService {
	return &CustomerService{
		Repo:&repository.CustomerDatabaseRepository{DB:pool.DatabaseConnection(b)},
	}
}