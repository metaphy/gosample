/*
Drawing
*/
package main

import (
    "fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

var (
	blue      color.Color = color.RGBA{0, 0, 255, 255}
	picwidth              = 640
	picheight             = 480
)

func main() {
    fmt.Println("Access localhost:9999")
	http.HandleFunc("/", handl)
	http.ListenAndServe(":9999", nil)
}

func handl(rw http.ResponseWriter, req *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, picwidth, picheight))
	// 填充蓝色,并把其写入到m
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	//以png编码格式,并将m写入到 rw里面去
	png.Encode(rw, m) //Encode writes the Image m to w in PNG format.

}
