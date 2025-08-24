package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	binary1()
}

func binary1() {
	sourceMat := gocv.IMRead("cmd/13_binary/book.jpg", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	threshold := gocv.Threshold(sourceMat, &targetMat, -1, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)
	fmt.Printf("%.2f\n", threshold)

	window1 := opencv.ShowWindow("binary1", sourceMat)
	defer window1.Close()
	window2 := opencv.ShowWindow("binary2", targetMat)
	defer window2.Close()

	window1.WaitKey(0)
	window2.WaitKey(0)

}
