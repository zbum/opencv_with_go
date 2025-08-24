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

	sourceMat := gocv.IMRead("cmd/17_detection/snowman.png", gocv.IMReadUnchanged)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	defer targetMat.Close()

	err := gocv.Canny(sourceMat, &targetMat, 150, 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	window1 := opencv.ShowWindow("detection1", sourceMat)
	defer window1.Close()

	window2 := opencv.ShowWindow("detection2", targetMat)
	defer window2.Close()
	window2.MoveWindow(100, 100)

	minValThreadHold := window2.CreateTrackbar("threshold1", 255)
	maxValThreadHold := window2.CreateTrackbar("threshold2", 255)

	for {
		maxVal := maxValThreadHold.GetPos()
		minVal := minValThreadHold.GetPos()

		err := gocv.Canny(sourceMat, &targetMat, float32(minVal), float32(maxVal))
		if err != nil {
			fmt.Println(err)
			return
		}

		opencv.ShowWindow("detection2", targetMat)

		if window2.WaitKey(1) == 'q' {
			break
		}
	}

}
