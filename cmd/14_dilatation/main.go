package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"opencv"
)

func main() {
	binary1()
}

func binary1() {
	sourceMat := gocv.IMRead("cmd/14_dilatation/dilate.png", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	defer targetMat.Close()

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{3, 3})
	defer kernel.Close()

	err := gocv.DilateWithParams(
		sourceMat,
		&targetMat,
		kernel,
		image.Point{X: -1, Y: -1},
		3,
		gocv.BorderConstant,
		color.RGBA{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	window1 := opencv.ShowWindow("binary1", targetMat)
	defer window1.Close()

	window1.WaitKey(0)

}
