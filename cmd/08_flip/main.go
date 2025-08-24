package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {
	flip(1)  //좌우 대칭
	flip(0)  //상하 대칭
	flip(-1) //상하좌우대칭

}

func flip(flipCode int) {
	imageMat := gocv.IMRead("cmd/08_flip/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.Flip(imageMat, &imageMat, flipCode)
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("flip-%d", flipCode), imageMat)
	defer window.Close()
	window.WaitKey(0)
}
