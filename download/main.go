package main

import (
	"fmt"
	"io"
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

// view the body and find all the links
func view(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = view(links, c)
	}
	return links
}

// get the Big Picture link
func bigImg(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, img := range n.Attr {
			if img.Key == "id" && img.Val == "bigImg" {
				for _, imgB := range n.Attr {
					if imgB.Key == "src" {
						imgURL := imgB.Val
						imgNameIndex := strings.LastIndex(imgURL, "/")
						imgName := imgURL[imgNameIndex+1:]
						// fmt.Println(imgB.Val)
						fmt.Printf("%s", imgName)
						//download
						save(imgName, imgURL)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		bigImg(c)
	}
}

// download the image
func save(file string, url string) {
	f, err := os.Create(file)
	check(err)
	defer f.Close()

	resp, err := http.Get(url)
	check(err)

	_, err = io.Copy(f, resp.Body)
	check(err)
	defer resp.Body.Close()
	fmt.Println("\tdone!")
}

// main function
func main() {
	site := "http://desk.zol.com.cn"
	resp, err := http.Get(site)
	check(err)

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	check(err)

	linksMap := make(map[string]bool)

	for _, link := range view(nil, doc) {
		if strings.Index(link, "bizhi") == 1 { //hard-coded
			linksMap[link] = false
		}
	}

	for link := range linksMap {
		url := fmt.Sprintf("%s%s", site, link)
		//fmt.Printf("%s\n", url)
		resp, err = http.Get(url)
		check(err)
		doc, err = html.Parse(resp.Body)
		resp.Body.Close()
		check(err)

		bigImg(doc)
	}
}
