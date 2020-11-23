package main

import (
	"bufio"
	"bytes"
	"image/jpeg"
	"os"
)

type Camera struct {

}

func NewCamera() *Camera {
	return &Camera{}
}

func (*Camera) GetImage() ([]byte, error) {
	file, err := os.Open("/Users/fponce/Downloads/sasha.jpg")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bufr := bufio.NewReader(file)

	image, err := jpeg.Decode(bufr)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	jpeg.Encode(writer, image, &jpeg.Options{Quality: 1})
	return b.Bytes(), err
}

