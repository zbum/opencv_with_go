package main

import (
	"fmt"
	"image"
	"image/color"
	"opencv"

	"gocv.io/x/gocv"
)

func main() {

	textWithEnglish()
	textWithKorean()

}

func textWithEnglish() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)

	scale := float64(1)
	whiteColor := color.RGBA{255, 255, 255, 255}
	thickness := 1

	err := gocv.PutText(&img, "Nado Simplex", image.Point{20, 50}, gocv.FontHersheySimplex, scale, whiteColor, thickness)
	if err != nil {
		return
	}

	err = gocv.PutText(&img, "Nado Plain", image.Point{20, 150}, gocv.FontHersheyPlain, scale, whiteColor, thickness)
	if err != nil {
		return
	}

	err = gocv.PutText(&img, "Nado Script Simplex", image.Point{20, 250}, gocv.FontHersheyScriptSimplex, scale, whiteColor, thickness)
	if err != nil {
		return
	}

	err = gocv.PutText(&img, "Nado Triplex", image.Point{20, 350}, gocv.FontHersheyTriplex, scale, whiteColor, thickness)
	if err != nil {
		return
	}

	err = gocv.PutText(&img, "Nado Italic", image.Point{20, 450}, gocv.FontHersheyTriplex|gocv.FontItalic, scale, whiteColor, thickness)
	if err != nil {
		return
	}

	window := opencv.ShowWindow("text", img)
	defer window.Close()
	window.WaitKey(0)
}

// 실패.. 일단 후퇴
// https://youtu.be/XK3eU9egll8?si=ybxDHL1Onh2Bo53Z&t=5002
func textWithKorean() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)

	// 한글 텍스트 표시 시도
	whiteColor := color.RGBA{255, 255, 255, 255}
	scale := float64(1)
	thickness := 1

	err := gocv.PutText(&img, "나도 코딩", image.Point{20, 50}, gocv.FontHersheySimplex, scale, whiteColor, thickness)
	if err != nil {
		fmt.Println(err)
		return
	}

	window := opencv.ShowWindow("Korean Text", img)
	defer window.Close()
	window.WaitKey(0)
}
