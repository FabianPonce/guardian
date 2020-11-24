package main

import (
	"errors"
	"github.com/blackjack/webcam"
)

type Camera struct {

}

func NewCamera() *Camera {
	return &Camera{}
}

func (*Camera) GetImage() ([]byte, error) {
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
		return frame, err
	}

	return nil, errors.New("No frame")
}

