package main

import (
	"bytes"
	"errors"
	"github.com/blackjack/webcam"
	"io/ioutil"
)

type CameraImpl struct {

}

func NewCamera() *CameraImpl {
	return &CameraImpl{}
}

func (*CameraImpl) GetImage() ([]byte, error) {
	webcam, err := webcam.Open("/dev/video0")
	if err != nil {
		return nil, err
	}
	defer webcam.Close()

	err = webcam.StartStreaming()
	if err != nil {
		return nil, err
	}

	err = webcam.WaitForFrame(1000)
	if err != nil {
		return nil, err
	}

	frame, err := webcam.ReadFrame()
	if err != nil {
		return nil, err
	}

	if len(frame) != 0 {
		var buf bytes.Buffer
		buf.Write(frame)

		ioutil.WriteFile("frame.jpg", buf.Bytes(), 644)

		return buf.Bytes(), err
	}

	return nil, errors.New("No frame")
}

