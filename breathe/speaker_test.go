package breathe

import (
	"testing"
)

func TestNewSpeaker(t *testing.T) {
	t.Run("no sound", func(t *testing.T) {
		speaker := NewSpeaker("none")
		if speaker.inhaleStreamer != nil {
			t.Fatalf("sound of 'none' should not init streamers")
		}
	})

	t.Run("some sound", func(t *testing.T) {
		speaker := NewSpeaker("some")
		if speaker.inhaleStreamer == nil {
			t.Fatalf("sound of not 'none' should init streamers")
		}
	})
}
