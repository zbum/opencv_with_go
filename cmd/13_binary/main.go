package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	//binary1()
	binary2()
}

func binary1() {
	sourceMat := gocv.IMRead("cmd/13_binary/book.jpg", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	gocv.Threshold(sourceMat, &targetMat, 127, 255, gocv.ThresholdBinary)

	window1 := opencv.ShowWindow("binary1", sourceMat)
	defer window1.Close()
	window2 := opencv.ShowWindow("binary2", targetMat)
	defer window2.Close()

	window1.WaitKey(0)
	window2.WaitKey(0)

}

func binary2() {
	sourceMat := gocv.IMRead("cmd/13_binary/book.jpg", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	gocv.Threshold(sourceMat, &targetMat, 127, 255, gocv.ThresholdBinary)

	window1 := gocv.NewWindow("binary3")
	defer window1.Close()

	trackbar := window1.CreateTrackbar("threshold", 255)
	trackbar.SetMin(0)
	trackbar.SetPos(127)
	
	err := window1.IMShow(targetMat)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		pos := trackbar.GetPos()
		gocv.Threshold(sourceMat, &targetMat, float32(pos), 255, gocv.ThresholdBinary)
		window1.IMShow(targetMat)

		if window1.WaitKey(1) == 'q' {
			break
		}
	}

	window1.WaitKey(0)

}
