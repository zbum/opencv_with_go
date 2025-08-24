package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	convert(gocv.ColorBGRToGray)
}

func convert(code gocv.ColorConversionCode) {
	imageMat := gocv.IMRead("cmd/10_convertImage/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.CvtColor(imageMat, &imageMat, code)
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("convert-%d", code), imageMat)
	defer window.Close()
	window.WaitKey(0)
}
