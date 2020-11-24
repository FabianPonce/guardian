package main

type Camera interface {
	GetImage() ([]byte, error)
}
