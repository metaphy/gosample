package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// error handler
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// forEachNode
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// ExactLinks exact the links of the url
func ExactLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("parsing error. URL %s:%s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing error. URL %s:%v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue //ignore the parse url error
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// BreadthFirst search
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// download the images
func crawl(url string) []string {
	// fmt.Println(url)
	// download
	resp, err := http.Get(url)
	logError(err)
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	logError(err)
	findImageNode(doc)

	list, err := ExactLinks(url)
	logError(err)
	return list
}

// check the error and continue
func logError(err error) {
	if err != nil {
		log.Print(err)
	}
}

// get the Big Picture link
func findImageNode(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, img := range n.Attr {
			if img.Key == "src" {
				save(img.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findImageNode(c)
	}
}

var saved = make(map[string]bool)

// download the image
func save(imgURL string) {
	fmt.Println(imgURL)
	if len(imgURL) < 5 {
		return
	}
	_, ok := saved[imgURL]
	if !ok {
		saved[imgURL] = true
	} else {
		return
	}

	fileType := string(imgURL[len(imgURL)-3:])
	if fileType != "jpg" && fileType != "png" && fileType != "gif" {
		return
	}
	imgNameIndex := strings.LastIndex(imgURL, "/")
	imgName := imgURL[imgNameIndex+1:]

	savePath := "/Users/peter/Pictures/d0/"
	file := fmt.Sprintf("%s%s", savePath, imgName)

	f, err := os.Create(file)
	check(err)
	defer f.Close()

	resp, err := http.Get(imgURL)
	if err == nil {
		_, err := io.Copy(f, resp.Body)
		check(err)
		defer resp.Body.Close()
		fmt.Println("done!")
	}
}

// main function
func main() {
	url := []string{
		// "http://desk.zol.com.cn/",
		"http://somesite",
	}
	breadthFirst(crawl, url)
}
