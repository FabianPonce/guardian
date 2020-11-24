package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

func Base64Encode(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

func main() {
	camera := NewCamera()
	classifier := NewClassifier()

	start := time.Now()

	image, err := camera.GetImage()
	if err != nil {
		panic(err)
	}

	isDog, err := classifier.ContainsDog(Base64Encode(image))
	if err != nil {
		panic(err)
	}

	if isDog {
		fmt.Println("There's a dog!")
	} else {
		fmt.Println("No dog.")
	}

	duration := time.Since(start)

	fmt.Printf("Classification complete in %vms\n", duration.Milliseconds())
}
