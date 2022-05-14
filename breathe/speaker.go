package breathe

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	beepSpeaker "github.com/faiface/beep/speaker"
)

type Speaker struct {
	sound          string
	format         beep.Format
	inhaleStreamer beep.StreamSeekCloser
	exhaleStreamer beep.StreamSeekCloser
	holdStreamer   beep.StreamSeekCloser
	oneStreamer    beep.StreamSeekCloser
	twoStreamer    beep.StreamSeekCloser
	threeStreamer  beep.StreamSeekCloser
	fourStreamer   beep.StreamSeekCloser
	fiveStreamer   beep.StreamSeekCloser
	sixStreamer    beep.StreamSeekCloser
	sevenStreamer  beep.StreamSeekCloser
	eightStreamer  beep.StreamSeekCloser
	nineStreamer   beep.StreamSeekCloser
}

// Returns a function to play the predefined sounds
// by specifying the soundName
func NewSpeaker(sound string) Speaker {
	speaker := Speaker{sound: sound}

	// Don't init the streamers, if there is no sound anyways
	if speaker.sound == "none" {
		return speaker
	}

	speaker.inhaleStreamer, speaker.format = initStreamer("inhale.mp3")
	speaker.exhaleStreamer, _ = initStreamer("exhale.mp3")
	speaker.holdStreamer, _ = initStreamer("hold.mp3")
	speaker.oneStreamer, _ = initStreamer("1.mp3")
	speaker.twoStreamer, _ = initStreamer("2.mp3")
	speaker.threeStreamer, _ = initStreamer("3.mp3")
	speaker.fourStreamer, _ = initStreamer("4.mp3")
	speaker.fiveStreamer, _ = initStreamer("5.mp3")
	speaker.sixStreamer, _ = initStreamer("6.mp3")
	speaker.sevenStreamer, _ = initStreamer("7.mp3")
	speaker.eightStreamer, _ = initStreamer("8.mp3")
	speaker.nineStreamer, _ = initStreamer("9.mp3")

	beepSpeaker.Init(speaker.format.SampleRate, speaker.format.SampleRate.N(time.Second/10))

	return speaker
}

func (speaker *Speaker) PlaySound(soundName string) {
	if speaker.sound == "none" {
		return
	}

	switch soundName {
	case "Inhale":
		if speaker.sound == "words" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.inhaleStreamer)
			defer speaker.inhaleStreamer.Seek(0)
		}
	case "Exhale":
		if speaker.sound == "words" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.exhaleStreamer)
			defer speaker.exhaleStreamer.Seek(0)
		}
	case "Hold":
		if speaker.sound == "words" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.holdStreamer)
			defer speaker.holdStreamer.Seek(0)
		}
	case "1":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.oneStreamer)
			defer speaker.oneStreamer.Seek(0)
		}
	case "2":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.twoStreamer)
			defer speaker.twoStreamer.Seek(0)
		}
	case "3":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.threeStreamer)
			defer speaker.threeStreamer.Seek(0)
		}
	case "4":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.fourStreamer)
			defer speaker.fourStreamer.Seek(0)
		}
	case "5":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.fiveStreamer)
			defer speaker.fiveStreamer.Seek(0)
		}
	case "6":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.sixStreamer)
			defer speaker.sixStreamer.Seek(0)
		}
	case "7":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.sevenStreamer)
			defer speaker.sevenStreamer.Seek(0)
		}
	case "8":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.eightStreamer)
			defer speaker.eightStreamer.Seek(0)
		}
	case "9":
		if speaker.sound == "numbers" || speaker.sound == "all" {
			beepSpeaker.Play(speaker.nineStreamer)
			defer speaker.nineStreamer.Seek(0)
		}
	}
}

//go:embed "assets/*.mp3"
var assets embed.FS

// Initializes a streamer for the given file under the assets folder
func initStreamer(fileName string) (beep.StreamSeekCloser, beep.Format) {
	inhaleFile, err := assets.Open(fmt.Sprintf("assets/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}

	inhaleStreamer, format, err := mp3.Decode(inhaleFile)
	if err != nil {
		log.Fatal(err)
	}

	return inhaleStreamer, format
}
