package breathe

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Returns a function to play the predefined sounds
// by specifying the soundName
func initSpeaker(sound string) func(soundName string) {
	inhaleStreamer, format := initStreamer("inhale.mp3")
	exhaleStreamer, _ := initStreamer("exhale.mp3")
	holdStreamer, _ := initStreamer("hold.mp3")
	oneStreamer, _ := initStreamer("1.mp3")
	twoStreamer, _ := initStreamer("2.mp3")
	threeStreamer, _ := initStreamer("3.mp3")
	fourStreamer, _ := initStreamer("4.mp3")
	fiveStreamer, _ := initStreamer("5.mp3")
	sixStreamer, _ := initStreamer("6.mp3")
	sevenStreamer, _ := initStreamer("7.mp3")
	eightStreamer, _ := initStreamer("8.mp3")
	nineStreamer, _ := initStreamer("9.mp3")

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return func(soundName string) {
		if sound == "none" {
			return
		}
		switch soundName {
		case "Inhale":
			if sound == "words" || sound == "all" {
				speaker.Play(inhaleStreamer)
				defer inhaleStreamer.Seek(0)
			}
		case "Exhale":
			if sound == "words" || sound == "all" {
				speaker.Play(exhaleStreamer)
				defer exhaleStreamer.Seek(0)
			}
		case "Hold":
			if sound == "words" || sound == "all" {
				speaker.Play(holdStreamer)
				defer holdStreamer.Seek(0)
			}
		case "1":
			if sound == "numbers" || sound == "all" {
				speaker.Play(oneStreamer)
				defer oneStreamer.Seek(0)
			}
		case "2":
			if sound == "numbers" || sound == "all" {
				speaker.Play(twoStreamer)
				defer twoStreamer.Seek(0)
			}
		case "3":
			if sound == "numbers" || sound == "all" {
				speaker.Play(threeStreamer)
				defer threeStreamer.Seek(0)
			}
		case "4":
			if sound == "numbers" || sound == "all" {
				speaker.Play(fourStreamer)
				defer fourStreamer.Seek(0)
			}
		case "5":
			if sound == "numbers" || sound == "all" {
				speaker.Play(fiveStreamer)
				defer fiveStreamer.Seek(0)
			}
		case "6":
			if sound == "numbers" || sound == "all" {
				speaker.Play(sixStreamer)
				defer sixStreamer.Seek(0)
			}
		case "7":
			if sound == "numbers" || sound == "all" {
				speaker.Play(sevenStreamer)
				defer sevenStreamer.Seek(0)
			}
		case "8":
			if sound == "numbers" || sound == "all" {
				speaker.Play(eightStreamer)
				defer eightStreamer.Seek(0)
			}
		case "9":
			if sound == "numbers" || sound == "all" {
				speaker.Play(nineStreamer)
				defer nineStreamer.Seek(0)
			}
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
