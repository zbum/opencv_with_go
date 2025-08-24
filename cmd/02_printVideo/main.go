package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
)

func main() {

	cap, err := gocv.VideoCaptureFile("cmd/02_printVideo/video.mp4")
	if err != nil {
		return
	}

	for cap.IsOpened() {
		frame := gocv.NewMat()
		success := cap.Read(&frame)
		if !success {
			fmt.Println("No more frame.")
			break
		}

		window := opencv.ShowWindow("video", frame)
		defer window.Close()

		if window.WaitKey(25) == 'q' {
			fmt.Println("Exit by user.")
			break
		}

	}
}
