package alerter

import (
	"github.com/padster/go-sound/output"
	"github.com/padster/go-sound/sounds"
)

type AudioAlerter struct {
	uri string
	playing bool
}

func NewAudioAlerter(uri string) *AudioAlerter {
	return &AudioAlerter{
		uri: uri,
		playing: false,
	}
}

func (a *AudioAlerter) Alert() error {
	wav := sounds.LoadWavAsSound(a.uri, 0)
	output.Play(wav)
	return nil
}