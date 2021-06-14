package proxy

func ExampleProxy() {
	proxy := NewProxy()

	//Output:
	//begin
	//action
	//end

	proxy.do()

}
