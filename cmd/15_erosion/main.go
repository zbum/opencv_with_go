package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"opencv"
)

func main() {
	binary1()
}

func binary1() {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{3, 3})
	defer kernel.Close()

	sourceMat := gocv.IMRead("cmd/15_erosion/erode.png", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	defer targetMat.Close()

	err := gocv.ErodeWithParams(
		sourceMat,
		&targetMat,
		kernel,
		image.Point{X: -1, Y: -1},
		3,
		int(gocv.BorderConstant),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	window1 := opencv.ShowWindow("binary1", targetMat)
	defer window1.Close()

	window1.WaitKey(0)

}
