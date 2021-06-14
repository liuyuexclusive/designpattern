package abstract_factory

import "fmt"

type OrderMainDao interface {
	SaveMainOrder()
}

type OrderDetailDao interface {
	SaveDetailOrder()
}

type Factory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

type RDBMainDao struct{}

func (r *RDBMainDao) SaveMainOrder() {
	fmt.Println("rdb save main order")
}

type RDBDetailDao struct{}

func (r *RDBDetailDao) SaveDetailOrder() {
	fmt.Println("rdb save detail order")
}

type XMLMainDao struct{}

func (x *XMLMainDao) SaveMainOrder() {
	fmt.Println("xml save main order")
}

type XMLDetailDao struct{}

func (x *XMLDetailDao) SaveDetailOrder() {
	fmt.Println("xml save detail order")
}

type RDBFactory struct{}

func (r *RDBFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}
func (r *RDBFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

type XMLFactory struct{}

func (r *XMLFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}
func (r *XMLFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}
