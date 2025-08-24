package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"opencv"
)

func main() {
	gaussianBlur(1)
	gaussianBlur(2)
	gaussianBlur(3)
}

func gaussianBlur(sigmaX float64) {
	imageMat := gocv.IMRead("cmd/11_blur/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.GaussianBlur(imageMat, &imageMat, image.Point{}, sigmaX, 0, gocv.BorderConstant)
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("blur-%f", sigmaX), imageMat)
	defer window.Close()
	window.WaitKey(0)
}
