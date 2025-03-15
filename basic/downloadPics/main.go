package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"
)

// main function
func main() {
	picUrl := "https://shuicaowang.com/wp-content/uploads/2019/03/20190329-CIAC-yulin-numbern.jpg"
	ch := make(chan string)
	
	for i := 6; i <= 11; i++ {
		// get the pic's specific URL
		thePicUrl := strings.Replace(picUrl, "numbern", fmt.Sprintf("%d", i) , 1)
		go fetchAndSave(thePicUrl, ch)
	}

	for i := 6; i <= 11; i++ {
		fmt.Println(<-ch)
	}
}

func fetchAndSave (thePicUrl string, ch chan <- string) {
	// save-to file
	sarr := strings.Split(thePicUrl, "/")
	fileName := sarr[len(sarr)-1]
	folder := "/home/peter/Pictures/tank/"

	f, err := os.Create(folder + fileName)
	if err != nil { 
		ch <- fmt.Sprint(err) // error sent to channel
		return
	}
	defer f.Close()

	resp, err := http.Get(thePicUrl)
	if err != nil { 
		ch <- fmt.Sprint(err) // error sent to channel
		return
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while copying %s:%v", thePicUrl, err)
		return
	}

	defer resp.Body.Close()
	ch <- fmt.Sprint(thePicUrl + "\n ... done!")
}
