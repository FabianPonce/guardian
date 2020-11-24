package main

import (
	"bytes"
	"github.com/utahta/go-openuri"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"io/ioutil"
	"time"
)

type AudioAlerter struct {
	uri string
}

func NewAudioAlerter(uri string) *AudioAlerter {
	return &AudioAlerter{
		uri: uri,
	}
}

func (a *AudioAlerter) Alert() error {
	f, err := openuri.Open(a.uri)
	if err != nil {
		return err
	}
	defer f.Close()

	// data must be copied into a new buffer otherwise it will be lost as the sound plays but this function exists
	// due to the deferred f.Close()
	dat, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(dat)

	streamer, format, err := wav.Decode(buf)
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