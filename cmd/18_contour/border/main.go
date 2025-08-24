package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
	"opencv"
)

func main() {
	binary1()
}

func binary1() {

	sourceMat := gocv.IMRead("cmd/18_contour/card.png", gocv.IMReadUnchanged)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	defer targetMat.Close()

	err := sourceMat.CopyTo(&targetMat)
	if err != nil {
		return
	}

	grayMat := gocv.NewMat()
	defer grayMat.Close()
	err = gocv.CvtColor(sourceMat, &grayMat, gocv.ColorBGRToGray)
	if err != nil {
		return
	}

	otsuMat := gocv.NewMat()
	defer otsuMat.Close()
	gocv.Threshold(
		grayMat,
		&otsuMat,
		-1,
		255,
		gocv.ThresholdBinary|gocv.ThresholdOtsu,
	)

	//contours := gocv.FindContours(otsuMat, gocv.RetrievalExternal, gocv.ChainApproxNone)
	contours := gocv.FindContours(otsuMat, gocv.RetrievalTree, gocv.ChainApproxNone)

	fmt.Println("contours:", contours.ToPoints())

	rgba := color.RGBA{G: 200}
	err = gocv.DrawContours(&targetMat, contours, -1, rgba, 2)
	if err != nil {
		return
	}

	window1 := opencv.ShowWindow("source", sourceMat)
	defer window1.Close()

	window2 := opencv.ShowWindow("gray", grayMat)
	defer window2.Close()
	window2.MoveWindow(100, 100)

	window3 := opencv.ShowWindow("otsu", otsuMat)
	defer window3.Close()
	window3.MoveWindow(200, 100)

	window4 := opencv.ShowWindow("contour", targetMat)
	defer window4.Close()
	window4.MoveWindow(300, 100)

	window4.WaitKey(0)

}
