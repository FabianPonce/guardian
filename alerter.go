package main

type Alerter interface {
	Alert() error
}
