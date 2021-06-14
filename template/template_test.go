package template

func ExampleTemplate() {
	//Output:
	//http download from uri www.baidu.com
	//http save
	//ftp download from uri www.google.com
	//ftp save
	t := NewTemplate(&HttpDownload{}, "www.baidu.com")

	t.download()

	t2 := NewTemplate(&FtpDownload{}, "www.google.com")

	t2.download()

}
