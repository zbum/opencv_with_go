package main

import (
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	unchangedMat := gocv.IMRead("cmd/01_printImage/img.png", gocv.IMReadUnchanged)
	defer unchangedMat.Close()

	grayMat := gocv.IMRead("cmd/01_printImage/img.png", gocv.IMReadGrayScale)
	defer grayMat.Close()

	colorMat := gocv.IMRead("cmd/01_printImage/img.png", gocv.IMReadColor)
	defer colorMat.Close()

	unchangedWindow := opencv.ShowWindow("unchanged", unchangedMat)
	defer unchangedWindow.Close()
	grayWindow := opencv.ShowWindow("gray", grayMat)
	defer grayWindow.Close()
	colorWindow := opencv.ShowWindow("color", colorMat)
	defer colorWindow.Close()

	unchangedWindow.WaitKey(0)
	grayWindow.WaitKey(0)
	colorWindow.WaitKey(0)

}
