package main

import "log"

const(
	GuardianConfigurationFile = "guardian.yml"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Printf("Unable to load configuration file %v.\n", GuardianConfigurationFile)
	}

	camera := config.CreateCamera()
	err = camera.Open()
	if err != nil {
		log.Println("Unable to initialize camera device for reading.")
		panic(err)
	}
	defer camera.Close()

	alerter, err := config.CreateAlerter()
	if err != nil {
		panic(err)
	}

	classifier, err := config.CreateClassifier()
	if err != nil {
		panic(err)
	}

	guardian := NewGuardian(camera, classifier, alerter, config.Guardian.Criteria, config.GetGuardianOptions())
	guardian.Run()
}
