package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"time"
	"github.com/corona10/goimagehash"
)

const(
	MinimumDistanceMotionDetectionThreshold = 15
)

type Guardian struct {
	Camera *CameraImpl
	Classifier *Classifier
	initialized bool

	lastHash *goimagehash.ImageHash
}

func NewGuardian() *Guardian {
	return &Guardian{
		Camera: NewCamera(),
		Classifier: NewClassifier(),
		initialized: false,
	}
}

func imageHashFromBytes(b []byte) (*goimagehash.ImageHash, error) {
	img, err := jpeg.Decode(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	hash, err := goimagehash.AverageHash(img)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (g *Guardian) Run() {
	for true {
		if !g.initialized {
			g.Camera.Configure()
			g.initialized = true
			defer g.Camera.Close()
		}

		img, err := g.Camera.GetImage()
		if err != nil {
			fmt.Println(err)
			continue
		}

		currentHash, _ := imageHashFromBytes(img)
		distance := 0
		if g.lastHash != nil {
			distance, _ = g.lastHash.Distance(currentHash)
		}

		if g.lastHash != nil && distance < MinimumDistanceMotionDetectionThreshold {
			continue
		}
		g.lastHash = currentHash

		fmt.Printf("Movement detected, image distance of %v\n", distance)

		start := time.Now()

		meetsCriteria, err := g.classify(img)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if meetsCriteria {
			fmt.Println("Found dog!")
		} else {
			fmt.Println("No dog.")
		}

		fmt.Printf("Classification done in %v ms.", time.Now().Sub(start).Milliseconds())
	}
}

func (g *Guardian) classify(image []byte) (bool, error) {
	return g.Classifier.ContainsDog(image)
}
