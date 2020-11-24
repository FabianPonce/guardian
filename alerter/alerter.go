package alerter

type Alerter interface {
	Alert() error
}
