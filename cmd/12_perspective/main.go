package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"opencv"
)

func main() {
	perspective()
}

func perspective() {
	imageMat := gocv.IMRead("cmd/12_perspective/newspaper.jpg", gocv.IMReadUnchanged)
	defer imageMat.Close()

	src := gocv.NewPointVectorFromPoints([]image.Point{
		{511, 352},
		{1008, 345},
		{1122, 584},
		{455, 594},
	})

	dst := gocv.NewPointVectorFromPoints([]image.Point{
		{0, 0},
		{640, 0},
		{640, 240},
		{0, 240},
	})

	finalDest := gocv.Zeros(240, 640, gocv.MatTypeCV8UC4)
	transform := gocv.GetPerspectiveTransform(src, dst)
	err := gocv.WarpPerspective(imageMat, &finalDest, transform, image.Point{X: 640, Y: 240})
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("perspective"), finalDest)
	defer window.Close()
	window.WaitKey(0)
}
