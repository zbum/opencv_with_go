package main

import (
	"gocv.io/x/gocv"
	"image"
	"opencv"
)

func main() {
	cropImage()
}

func cropImage() {
	imageMat := gocv.IMRead("cmd/07_cropImage/img.png", gocv.IMReadUnchanged)
	defer imageMat.Close()

	croppedMat := imageMat.Region(image.Rectangle{Min: image.Point{X: 500, Y: 200}, Max: image.Point{X: 700, Y: 400}})
	defer croppedMat.Close()

	targetMat := imageMat.Region(image.Rectangle{Min: image.Point{X: 300, Y: 200}, Max: image.Point{X: 500, Y: 400}})
	for inx := 0; inx < targetMat.Rows(); inx++ {
		for jnx := 0; jnx < targetMat.Cols(); jnx++ {
			targetMat.SetFloatAt(inx, jnx, croppedMat.GetFloatAt(inx, jnx))
		}
	}

	window := opencv.ShowWindow("resizeImageWithRatio", imageMat)
	defer window.Close()
	window.IMShow(imageMat)

	window.WaitKey(0)
}
