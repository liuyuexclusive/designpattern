package template

import (
	"fmt"
	"unsafe"
)

type Downloader interface {
	download()
}

type implement interface {
	download(string)
	save()
}

type template struct {
	implement implement
	uri       string
}

func NewTemplate(implement implement, uri string) *template {
	return &template{implement: implement, uri: uri}
}

func (t *template) download() {
	t.implement.download(t.uri)
	t.implement.save()
}

type HttpDownload struct{}

func (h *HttpDownload) download(uri string) {
	fmt.Printf("http download from uri %s\n", uri)
}
func (h *HttpDownload) save() {
	fmt.Printf("http save\n")
}

type FtpDownload struct{}

func (f *FtpDownload) download(uri string) {
	fmt.Printf("ftp download from uri %s\n", uri)
}

func (h *FtpDownload) save() {
	fmt.Printf("ftp save\n")
}

func test() {
	var a int
	b := unsafe.Sizeof(a)
	fmt.Println(b)
}
