package main

import (
	"fmt"
	"time"
)

type Guardian struct {
	Camera *CameraImpl
	Classifier *Classifier
	initialized bool
}

func NewGuardian() *Guardian {
	return &Guardian{
		Camera: NewCamera(),
		Classifier: NewClassifier(),
		initialized: false,
	}
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
