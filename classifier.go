package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type Classifier struct {
	session *session.Session
	client *rekognition.Rekognition
}

func NewClassifier() *Classifier {
	s := session.New()
	return &Classifier{
		session: s,
		client: rekognition.New(s),
	}
}

func (c *Classifier) getLabels(image *rekognition.Image) ([]*rekognition.Label, error) {
	input := rekognition.DetectLabelsInput{
		Image:         image,
		MaxLabels:     aws.Int64(64),
		MinConfidence: aws.Float64(50),
	}

	response, err := c.client.DetectLabels(&input)
	if err != nil {
		return nil, err
	}

	return response.Labels, nil
}

func (c *Classifier) ContainsDog(image []byte) (bool, error) {
	labels, error := c.getLabels(&rekognition.Image{
		Bytes:    image,
		S3Object: nil,
	})

	if error != nil {
		return false, error
	}

	for _, label := range labels {
		n := *label.Name
		fmt.Printf("Label: %v | Probability: %v\n", n, *label.Confidence)
		if 	n == "Dog" ||
			n == "Canine" ||
			n == "Pet" ||
			n == "Animal" {
			return true, nil
		}
	}

	return false, nil
}