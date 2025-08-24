package opencv

import (
	"fmt"
	"gocv.io/x/gocv"
)

func ShowWindow(title string, inputMat gocv.Mat) *gocv.Window {
	window := gocv.NewWindow(title)
	err := window.IMShow(inputMat)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return window
}
