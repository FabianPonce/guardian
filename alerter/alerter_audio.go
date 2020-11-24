package alerter

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/utahta/go-openuri"
	"io/ioutil"
	"time"
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
	if a.playing {
		fmt.Println("Skipping alert due to already playing sound.")
		return nil
	}

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

	a.playing = true
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		a.playing = false
	})))
	return nil
}