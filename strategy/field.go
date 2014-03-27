package main

import (
	"flag"
	"code.google.com/p/codewars/strategy/field"
	"image"
	"net/http"
	"fmt"
	"bytes"
	"encoding/base64"
	"image/png"
)

var maxPiek = flag.Int("maxPiek", 100, "max count for pieks")
var minPiek = flag.Int("minPiek", 50, "min count for pieks")
var size = flag.Int("size", 50, "field Size")

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head></head><body><img src=\"data:image/png;base64,"+show()+"\"></body></html>")
}

func show() string {		
	data := field.CreateField()
	m := image.NewNRGBA(image.Rect(0, 0, field.GetFieldSize(), field.GetFieldSize()))
	for y := 0; y < field.GetFieldSize(); y++ {
		for x := 0; x < field.GetFieldSize(); x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			if v == 0{
				//water
				m.Pix[i] = 0
				m.Pix[i+1] = 0
				m.Pix[i+2] = 255
			}else if v < 4{
				m.Pix[i] = 255
				m.Pix[i+1] = 172
				m.Pix[i+2] = 56
			}else{
				m.Pix[i] = 12
				m.Pix[i+1] = uint8(255-v)
				m.Pix[i+2] = 16
			}
			m.Pix[i+3] = 255
		}
	}
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func main() {
	flag.Parse()
	field.ChangeFieldSize(*size)
	field.MinPiek = *minPiek
	field.MaxPiek = *maxPiek
	http.HandleFunc("/field",ServeHTTP)
	http.ListenAndServe(":4000", nil)
	
	
	//for y:= range myField{
	//	fmt.Println(myField[y])
	//}
}
