package main

import (
	"errors"
	"github.com/creasty/defaults"
	"github.com/fabianponce/guardian/alerter"
	"github.com/fabianponce/guardian/camera"
	"github.com/fabianponce/guardian/classifier"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Camera struct {
		DeviceIndex int `yaml:"device_index"`
	} `yaml:"camera"`

	Drivers struct {
		Alert string `yaml:"alert"`
		Classification string `yaml:"classification"`
	} `yaml:"drivers"`

	Classification struct {
		Rekognition struct {
			MaxLabels int64 `default:"128" yaml:"max_labels"`
			MinConfidence float64 `default:"1" yaml:"min_confidence"`
		} `yaml:"rekognition"`
	} `yaml:"classification"`

	Alert struct {
		Sound struct {
			URI string `yaml:"uri"`
		} `yaml:"sound"`
	} `yaml:"alert"`

	Guardian struct {
		MotionThreshold int `yaml:"motion_threshold"`
		Criteria []ConfigMatchingCriteria `yaml:"criteria"`
	} `yaml:"guardian"`
}

type ConfigMatchingCriteria struct {
	Label string `yaml:"label"`
	Threshold float64 `yaml:"threshold"`
}

func LoadConfig() (Config, error) {
	config := Config{}

	contents, err := ioutil.ReadFile(GuardianConfigurationFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		return config, err
	}

	defaults.Set(config)

	return config, nil
}

func (c Config) CreateAlerter() (alerter.Alerter, error) {
	if c.Drivers.Alert == "sound" {
		return alerter.NewAudioAlerter(c.Alert.Sound.URI), nil
	} else {
		return nil, errors.New("unsupported alert driver")
	}
}

func (c Config) CreateCamera() camera.Camera {
	return camera.NewCamera(camera.CameraOptions{DeviceIndex: c.Camera.DeviceIndex})
}

func (c Config) CreateClassifier() (classifier.Classifier, error) {
	if c.Drivers.Classification == "rekognition" {
		return classifier.NewRekognitionClassifier(classifier.RekognitionOptions{
			MinConfidence: c.Classification.Rekognition.MinConfidence,
			MaxLabels:     c.Classification.Rekognition.MaxLabels,
		}), nil
	} else {
		return nil, errors.New("unsupported classification driver")
	}
}

func (c Config) GetGuardianOptions() GuardianOptions {
	return GuardianOptions{
		MotionDetectionThreshold: c.Guardian.MotionThreshold,
	}
}