package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	rotate(gocv.Rotate90Clockwise)        //시계방향 90 도 회전
	rotate(gocv.Rotate180Clockwise)       //180도 회전
	rotate(gocv.Rotate90CounterClockwise) //반시계반대방향으로 90도 회전

}

func rotate(rotateFlag gocv.RotateFlag) {
	imageMat := gocv.IMRead("cmd/09_rotate/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.Rotate(imageMat, &imageMat, rotateFlag)
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("rotate-%d", rotateFlag), imageMat)
	defer window.Close()
	window.WaitKey(0)
}
