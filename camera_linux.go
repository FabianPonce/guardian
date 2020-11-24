package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/blackjack/webcam"
)

type CameraImpl struct {
	webcam *webcam.Webcam
	options CameraOptions
}

func NewCamera(options CameraOptions) Camera {
	return &CameraImpl{
		options: options,
	}
}

func (c *CameraImpl) Open() error {
	webcam, err := webcam.Open(fmt.Sprintf("/dev/video%v", c.options.DeviceIndex))
	if err != nil {
		return err
	}
	c.webcam = webcam

	err = webcam.StartStreaming()
	return err
}

func (c *CameraImpl) Close() error {
	return c.webcam.Close()
}

func (c *CameraImpl) GetImage() ([]byte, error) {
	err := c.webcam.WaitForFrame(1000)
	if err != nil {
		return nil, err
	}

	frame, err := c.webcam.ReadFrame()
	if err != nil {
		return nil, err
	}

	if len(frame) != 0 {
		var buf bytes.Buffer
		buf.Write(frame)

		return buf.Bytes(), err
	}

	return nil, errors.New("No frame")
}

