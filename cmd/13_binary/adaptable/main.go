package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	binary2()
}

func binary2() {
	sourceMat := gocv.IMRead("cmd/13_binary/book.jpg", gocv.IMReadGrayScale)
	defer sourceMat.Close()

	targetMat := gocv.NewMat()
	gocv.Threshold(sourceMat, &targetMat, 127, 255, gocv.ThresholdBinary)

	window1 := gocv.NewWindow("binary3")
	defer window1.Close()

	blockSizeTrackbar := window1.CreateTrackbar("block_size", 100) //홀수만 가능
	blockSizeTrackbar.SetMin(1)
	blockSizeTrackbar.SetPos(25)

	c := window1.CreateTrackbar("c", 10)
	c.SetMin(1)
	c.SetPos(3)

	err := window1.IMShow(targetMat)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		blockSizePosition := blockSizeTrackbar.GetPos()
		cPosition := c.GetPos()

		if blockSizePosition <= 1 {
			blockSizePosition = 3
		}
		if blockSizePosition%2 == 0 {
			blockSizePosition += 1
		}

		gocv.AdaptiveThreshold(sourceMat, &targetMat, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, blockSizePosition, float32(cPosition))
		window1.IMShow(targetMat)

		if window1.WaitKey(1) == 'q' {
			break
		}
	}

	window1.WaitKey(0)

}
