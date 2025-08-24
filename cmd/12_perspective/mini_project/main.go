package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"opencv"
)

func main() {

	imageMat := gocv.IMRead("cmd/12_perspective/mini_project/img.jpg", gocv.IMReadUnchanged)
	defer imageMat.Close()

	window := opencv.ShowWindow("img", imageMat)
	defer window.Close()

	pointList := make([]image.Point, 0)
	userData := userData{*window, &pointList, imageMat}

	window.SetMouseHandler(mouseHandler, userData)
	window.WaitKey(0)

}

type userData struct {
	window    gocv.Window
	pointList *[]image.Point
	imageMat  gocv.Mat
}

func mouseHandler(event int, x int, y int, flags int, userdata interface{}) {
	switch event {
	case 0:
		return
	case 1:
		if castedUserData, ok := userdata.(userData); ok {
			*castedUserData.pointList = append(*castedUserData.pointList, image.Point{X: x, Y: y})
		}
	case 4:
		fmt.Println("mouse up")
	}

	refresh(userdata)

}

func refresh(userdata interface{}) {
	dotColor := color.RGBA{R: 255, B: 255, A: 255}

	if castedUserData, ok := userdata.(userData); ok {
		for _, point := range *castedUserData.pointList {
			err := gocv.Circle(&castedUserData.imageMat, point, 15, dotColor, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		castedUserData.window.IMShow(castedUserData.imageMat)

		if len(*castedUserData.pointList) == 4 {
			showResult(castedUserData.imageMat, castedUserData.pointList)
		}
	}
}

func showResult(imageMat gocv.Mat, pointList *[]image.Point) {
	width, height := 530, 710

	src := gocv.NewPointVectorFromPoints(*pointList)

	dst := gocv.NewPointVectorFromPoints([]image.Point{
		{0, 0},
		{width, 0},
		{width, height},
		{0, height},
	})

	finalDest := gocv.Zeros(height, width, gocv.MatTypeCV8UC4)
	transform := gocv.GetPerspectiveTransform(src, dst)
	err := gocv.WarpPerspective(imageMat, &finalDest, transform, image.Point{X: width, Y: height})
	if err != nil {
		return
	}

	window := opencv.ShowWindow(fmt.Sprintf("perspective"), finalDest)
	defer window.Close()
	window.WaitKey(0)

}
