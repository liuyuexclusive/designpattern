package abstract_factory

func Example_abstract_factory() {
	//Output:
	//rdb save main order
	//rdb save detail order
	//xml save main order
	//xml save detail order
	f1 := &RDBFactory{}
	f1.CreateOrderMainDao().SaveMainOrder()
	f1.CreateOrderDetailDao().SaveDetailOrder()
	f2 := &XMLFactory{}
	f2.CreateOrderMainDao().SaveMainOrder()
	f2.CreateOrderDetailDao().SaveDetailOrder()
}
