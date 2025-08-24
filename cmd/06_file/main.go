package main

import (
	"fmt"
	"image"
	"opencv"

	"gocv.io/x/gocv"
)

func main() {
	saveGrayImage()
	changeFormat()
	saveMovie()
	saveImageWithFixedSize()
	saveImageWithRatio()
	//보간법(interpolation)
	saveImageWithInterpolation()

}

func saveGrayImage() {
	gray := gocv.IMRead("cmd/06_file/img.png", gocv.IMReadGrayScale)

	window := opencv.ShowWindow("img", gray)
	defer window.Close()
	window.WaitKey(0)

	ok := gocv.IMWrite("cmd/06_file/img-gray.png", gray)
	if ok {
		fmt.Println("ok")
	}
}

func changeFormat() {
	gray := gocv.IMRead("cmd/06_file/img.png", gocv.IMReadGrayScale)
	ok := gocv.IMWrite("cmd/06_file/img-gray.jpg", gray)
	if ok {
		fmt.Println("ok")
	}
}

func saveMovie() {
	capture, err := gocv.VideoCaptureFile("cmd/06_file/video.mp4")
	if err != nil {
		fmt.Println("Error opening video file:", err)
		return
	}
	defer capture.Close()

	width := capture.Get(gocv.VideoCaptureFrameWidth)
	height := capture.Get(gocv.VideoCaptureFrameHeight)
	fps := capture.Get(gocv.VideoCaptureFPS)

	outputFile, err := gocv.VideoWriterFile("cmd/06_file/video-result.avi", "MJPG", fps, int(width), int(height), true)
	if err != nil {
		fmt.Println("Error creating video writer:", err)
		return
	}
	defer outputFile.Close()

	mat := gocv.NewMat()
	defer mat.Close()

	window := gocv.NewWindow("video")
	defer window.Close()

	fmt.Println("Processing video... Press 'q' to quit")

	for capture.IsOpened() {
		ok := capture.Read(&mat)
		if !ok {
			break
		}

		// Write the frame to the output video file
		outputFile.Write(mat)

		window.IMShow(mat)
		key := window.WaitKey(25)
		if key == 'q' {
			fmt.Println("User requested exit")
			break
		}
	}

	fmt.Println("Video processing completed. Output saved as video-result.avi")
}

func saveImageWithFixedSize() {
	imageMat := gocv.IMRead("cmd/06_file/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.Resize(imageMat, &imageMat, image.Point{400, 500}, 0, 0, gocv.InterpolationDefault)
	if err != nil {
		fmt.Println("Error resizing image:", err)
		return
	}

	window := opencv.ShowWindow("resizeImage", imageMat)
	defer window.Close()

	window.WaitKey(0)

}

func saveImageWithRatio() {
	imageMat := gocv.IMRead("cmd/06_file/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.Resize(imageMat, &imageMat, image.Point{}, 0.5, 0.5, gocv.InterpolationDefault)
	if err != nil {
		fmt.Println("Error resizing image:", err)
		return
	}

	window := opencv.ShowWindow("resizeImageWithRatio", imageMat)
	defer window.Close()

	window.WaitKey(0)
}

func saveImageWithInterpolation() {
	imageMat := gocv.IMRead("cmd/06_file/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	err := gocv.Resize(imageMat, &imageMat, image.Point{}, 1.5, 1.5, gocv.InterpolationCubic)
	if err != nil {
		fmt.Println("Error resizing image:", err)
		return
	}

	window := opencv.ShowWindow("resizeImageWithInterpolation", imageMat)
	defer window.Close()

	window.WaitKey(0)
}
