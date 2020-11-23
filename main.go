package main

import (
	"fmt"
	"time"
)

func main() {
	camera := NewCamera()
	classifier := NewClassifier()

	start := time.Now()

	image, err := camera.GetImage()
	if err != nil {
		panic(err)
	}

	isDog, err := classifier.ContainsDog(image)
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
