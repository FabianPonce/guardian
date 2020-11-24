package main

import (
	"bytes"
	"gocv.io/x/gocv"
	"image/jpeg"
)

type CameraImpl struct {

}

func NewCamera() *CameraImpl {
	return &CameraImpl{}
}

func (*CameraImpl) GetImage() ([]byte, error) {
	webcam, err := gocv.OpenVideoCapture(1)
	if err != nil {
		return nil, err
	}

	img := gocv.NewMat()
	webcam.Read(&img)

	image, err := img.ToImage()
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, image, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

