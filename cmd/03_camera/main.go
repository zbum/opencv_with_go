package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"opencv"
	"os"
)

func main() {
	cap, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", err)
		os.Exit(1)
	}

	for {
		frame := gocv.NewMat()
		ok := cap.Read(&frame)
		if !ok {
			break
		}

		window := opencv.ShowWindow("camera", frame)
		defer window.Close()
		if window.WaitKey(1) == 'q' {
			break
		}
	}

}
