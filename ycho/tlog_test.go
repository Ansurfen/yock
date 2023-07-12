package ycho

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func download(src, dst string) {
	res, err := http.Get(src)
	if err != nil {
		panic(err)
	}
	fp, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fp, Progress(res.ContentLength, res.Body))
	if err != nil {
		panic(err)
	}
}

func TestTlog(t *testing.T) {
	tlog, err := NewTLog(YchoOpt{})
	if err != nil {
		panic(err)
	}
	Set(tlog)
	go Eventloop()
	go func() {
		i := 0
		for {
			time.Sleep(2 * time.Second)
			Infof("%d: Hello World", i)
			i++
		}
	}()
	go download("https://download.oracle.com/java/20/latest/jdk-20_linux-aarch64_bin.tar.gz", "file1.zip")
	download("https://download.oracle.com/java/20/latest/jdk-20_linux-aarch64_bin.tar.gz", "file2.zip")
}
