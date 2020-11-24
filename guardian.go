package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"strings"
	"time"
	"github.com/corona10/goimagehash"
)

type Guardian struct {
	camera      Camera
	classifier  Classifier
	initialized bool
	alerter     Alerter
	criteria    []ConfigMatchingCriteria
	lastHash    *goimagehash.ImageHash
	options     GuardianOptions
}

type GuardianOptions struct {
	MotionDetectionThreshold int
}

func NewGuardian(camera Camera, classifier Classifier, alerter Alerter, criteria []ConfigMatchingCriteria, options GuardianOptions) *Guardian {
	return &Guardian{
		camera:      camera,
		classifier:  classifier,
		initialized: false,
		alerter:     alerter,
		criteria:    criteria,
		options: 	 options,
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
		img, err := g.camera.GetImage()
		if err != nil {
			fmt.Println(err)
			continue
		}

		currentHash, _ := imageHashFromBytes(img)
		distance := 0
		if g.lastHash != nil {
			distance, _ = g.lastHash.Distance(currentHash)
		}

		if g.lastHash != nil && distance < g.options.MotionDetectionThreshold {
			continue
		}
		g.lastHash = currentHash

		fmt.Printf("Movement detected, image distance of %v\n", distance)

		start := time.Now()

		labels, err := g.classifier.GetLabels(img)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.meetsCriteria(labels) {
			fmt.Println("Search criteria met, triggering alert.")
			err := g.alerter.Alert()
			if err != nil {
				fmt.Println("Unable to alert.")
				fmt.Println(err)
			}
		} else {
			fmt.Println("No criteria matched.")
		}

		fmt.Printf("Classification done in %v ms.\n", time.Now().Sub(start).Milliseconds())
	}
}

func (g *Guardian) meetsCriteria(labels []MatchedLabel) bool {
	for _, search := range g.criteria {
		for _, label := range labels {
			if strings.Compare(strings.ToLower(search.Label), strings.ToLower(label.Label)) == 0 &&
				label.Confidence >= search.Threshold {
				return true
			}
		}
	}

	return false
}
