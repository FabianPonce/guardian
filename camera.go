package main

import (
	"errors"
	"fmt"
	"github.com/blackjack/webcam"
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
		/*buf := &bytes.Buffer{}
		if err := jpeg.Encode(buf, img, nil); err != nil {
			log.Fatal(err)
			return
		}*/
		os.Stdout.Write(frame)
		os.Stdout.Sync()
		return frame, err
	}

	return nil, errors.New("No frame")
}

