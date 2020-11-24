package main

import (
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

type AudioAlerter struct {
	path string
}

func NewAudioAlerter(path string) *AudioAlerter {
	return &AudioAlerter{
		path: path,
	}
}

func (a *AudioAlerter) Alert() error {
	f, err := os.Open(a.path)
	if err != nil {
		return err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return err
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	speaker.Play(streamer)
	return nil
}