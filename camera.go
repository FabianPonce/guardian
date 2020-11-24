package main

type Camera interface {
	Configure() error
	GetImage() ([]byte, error)
	Close() error
}
