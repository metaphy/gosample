/*
Access the URLs list concurrently
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://baidu.com",
		"http://sohu.com",
		"http://taobao.com",
		"http://tmall.com",
		"http://qq.com",
		"http://sina.com.cn",
		"http://sm.cn",
		"http://360.cn",
		"http://jd.com",
		"http://alipay.com",
		"http://list.tmall.com",
		"http://weibo.com",
		"http://csdn.net",
		"http://samsung.com",
		"http://sogou.com",
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("Totally %.2fs eclapsed.\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // error sent to channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Don't leak memory
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, nbytes, url)
}
