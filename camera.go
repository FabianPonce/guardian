package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/blackjack/webcam"
	"image"
	"image/jpeg"
	"log"
	"os"
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

	fmap := webcam.GetSupportedFormats()
	fmt.Println("Available Formats: ")
	for p, s := range fmap {
		var pix []byte
		for i := 0; i < 4; i++ {
			pix = append(pix, byte(p>>uint(i*8)))
		}
		fmt.Printf("ID:%08x ('%s') %s\n   ", p, pix, s)
		for _, fs := range webcam.GetSupportedFrameSizes(p) {
			fmt.Printf(" %s", fs.GetString())
		}
		fmt.Printf("\n")
	}

	cmap := webcam.GetControls()
	fmt.Println("Available controls: ")
	for id, c := range cmap {
		fmt.Printf("ID:%08x %-32s  Min: %4d  Max: %5d\n", id, c.Name, c.Min, c.Max)
	}

	webcam.SetImageFormat(0x4745504a, 3280, 2464)

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
		img := image.NewYCbCr(image.Rect(0, 0, 3280, 2464), image.YCbCrSubsampleRatio422)
		for i := range img.Cb {
			ii := i * 4
			img.Y[i*2] = frame[ii]
			img.Y[i*2+1] = frame[ii+2]
			img.Cb[i] = frame[ii+1]
			img.Cr[i] = frame[ii+3]
		}
		buf := &bytes.Buffer{}
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		os.Stdout.Write(frame)
		os.Stdout.Sync()
		return frame, err
	}

	return nil, errors.New("No frame")
}

