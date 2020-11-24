package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type RekognitionClassifier struct {
	session *session.Session
	client *rekognition.Rekognition
	options RekognitionOptions
}

type RekognitionOptions struct {
	MinConfidence float64
	MaxLabels	int64
}

func NewRekognitionClassifier(options RekognitionOptions) Classifier {
	s := session.New()
	return &RekognitionClassifier{
		session: s,
		client: rekognition.New(s),
		options: options,
	}
}

func (c *RekognitionClassifier) GetLabels(image []byte) ([]MatchedLabel, error) {
	img := rekognition.Image{Bytes: image}
	input := rekognition.DetectLabelsInput{
		Image:         &img,
		MaxLabels:     aws.Int64(c.options.MaxLabels),
		MinConfidence: aws.Float64(c.options.MinConfidence),
	}

	response, err := c.client.DetectLabels(&input)
	if err != nil {
		return nil, err
	}

	var matches []MatchedLabel
	for _, label := range response.Labels {
		l := MatchedLabel{
			Label: *label.Name,
			Confidence: *label.Confidence,
		}
		matches = append(matches, l)
	}

	return matches, nil
}