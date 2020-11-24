package main

import (
	"bytes"
	"gocv.io/x/gocv"
	"image/jpeg"
)

type CameraImpl struct {
	webcam *gocv.VideoCapture
}

func NewCamera() *CameraImpl {
	return &CameraImpl{}
}

func (c *CameraImpl) Configure() error {
	webcam, err := gocv.OpenVideoCapture(1)
	if err != nil {
		return err
	}

	c.webcam = webcam
	return nil
}

func (c *CameraImpl) Close() error {
	return c.Close()
}

func (c *CameraImpl) GetImage() ([]byte, error) {
	img := gocv.NewMat()
	c.webcam.Read(&img)

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

