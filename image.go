package main

import (
	"bytes"
	"image/jpeg"
)
func reduceImageSize(image []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(image))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 1})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

