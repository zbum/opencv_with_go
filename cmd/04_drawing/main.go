package main

import (
	"image"
	"image/color"
	"opencv"

	"gocv.io/x/gocv"
)

func main() {
	rectangle()
	line()
	circle()
	polygon()
	textWithKorean()
}

func rectangle() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)
	defer img.Close()

	// 흰색 사각형 그리기: (100,100)에서 (200,200)까지
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255} // 흰색 (RGBA)
	rect := image.Rect(100, 100, 200, 200)              // (x1, y1, x2, y2)
	err := gocv.Rectangle(&img, rect, white, 3)
	if err != nil {
		return
	}

	window := opencv.ShowWindow("rectangle", img)
	defer window.Close()
	window.WaitKey(0)
}

func line() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)
	defer img.Close()

	color := color.RGBA{R: 255, G: 255, B: 100, A: 255}
	err := gocv.Line(&img, image.Point{X: 50, Y: 100}, image.Point{X: 400, Y: 50}, color, 3)
	if err != nil {
		return
	}

	window := opencv.ShowWindow("line", img)
	defer window.Close()
	window.WaitKey(0)

}

func circle() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)
	defer img.Close()

	white := color.RGBA{R: 255, G: 255, B: 0, A: 255} // 흰색 (RGBA)
	err := gocv.Circle(&img, image.Point{X: 200, Y: 100}, 50, white, -1)
	if err != nil {
		return
	}
	err = gocv.Circle(&img, image.Point{X: 400, Y: 100}, 50, white, 5)
	if err != nil {
		return
	}

	window := opencv.ShowWindow("circle", img)
	defer window.Close()
	window.WaitKey(0)
}

func polygon() {
	img := gocv.Zeros(480, 640, gocv.MatTypeCV8UC4)
	defer img.Close()

	// 빨간색 삼각형 그리기
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255} // 빨간색 (RGBA)

	points1 := []image.Point{
		{100, 100},
		{200, 100},
		{100, 200},
	}

	points2 := []image.Point{
		{200, 100},
		{300, 100},
		{300, 200},
	}

	err := gocv.Polylines(&img, gocv.NewPointsVectorFromPoints([][]image.Point{points1, points2}), true, red, 3)
	if err != nil {
		return
	}

	points3 := []image.Point{
		{100, 300},
		{200, 300},
		{100, 400},
	}

	points4 := []image.Point{
		{200, 300},
		{300, 300},
		{300, 400},
	}

	err = gocv.FillPoly(&img, gocv.NewPointsVectorFromPoints([][]image.Point{points3, points4}), red)
	if err != nil {
		return
	}

	window := opencv.ShowWindow("polygon", img)
	defer window.Close()
	window.WaitKey(0)
}
