package main

type Camera interface {
	Open() error
	GetImage() ([]byte, error)
	Close() error
}

type CameraOptions struct {
	DeviceIndex int
}