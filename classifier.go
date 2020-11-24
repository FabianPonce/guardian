package main

type MatchedLabel struct {
	Label string `yaml:"label"`
	Confidence float64 `yaml:"threshold"`
}

type Classifier interface {
	GetLabels([]byte) ([]MatchedLabel, error)
}

