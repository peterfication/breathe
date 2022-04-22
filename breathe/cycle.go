package breathe

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// A BreatheCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreatheCycle struct {
	Inhale     time.Duration
	InhaleHold time.Duration
	Exhale     time.Duration
	ExhaleHold time.Duration
}

func GenerateBreatheCycles(cycle BreatheCycle, cyclesCount int) (cycles []BreatheCycle) {
	for i := 0; i < cyclesCount; i++ {
		cycles = append(cycles, cycle)
	}
	return cycles
}

// Draw the UI and kick off the breath cycles
func RunBreatheCycles(title string, cycles []BreatheCycle, sound string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize TermUI: %v", err)
	}
	defer ui.Close()

	gaugeChart := initGaugeChart()
	textBox := initTextBox()
	renderText := createRenderText(textBox, title, cycles)
	playSound := initSpeaker(sound)

	go runBreatheCycles(cycles, gaugeChart, renderText, playSound)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

// The gauge chart that displays the breathing action to perform
func initGaugeChart() *widgets.Gauge {
	termWidth, termHeight := ui.TerminalDimensions()
	gaugeChart := widgets.NewGauge()
	gaugeChart.SetRect(0, 0, termWidth, termHeight-10)
	gaugeChart.Percent = 0
	gaugeChart.BarColor = ui.ColorGreen
	gaugeChart.BorderStyle.Fg = ui.ColorWhite
	gaugeChart.TitleStyle.Fg = ui.ColorCyan

	ui.Render(gaugeChart)
	return gaugeChart
}

// The text box that holds additional information about the breathing cycles
func initTextBox() *widgets.Paragraph {
	termWidth, termHeight := ui.TerminalDimensions()
	paragraph := widgets.NewParagraph()
	paragraph.SetRect(0, termHeight-10, termWidth, termHeight)
	paragraph.Text = "Breathe"

	ui.Render(paragraph)
	return paragraph
}

// The closure that refreshes the text box by re-rendering it
func createRenderText(textBox *widgets.Paragraph, title string, cycles []BreatheCycle) func(currentCycleCount int) {
	return func(currentCycleCount int) {
		text := fmt.Sprintf(`Always inhale through the nose!

%s
Total duration: %s
Cycle %d of %d
`, title, totalDuration(cycles), currentCycleCount, len(cycles))
		textBox.Text = text
		ui.Render(textBox)
	}
}

// Sum up the duration of all steps in all cycles
func totalDuration(cycles []BreatheCycle) time.Duration {
	totalDurationMilliseconds := int64(0)
	for _, cycle := range cycles {
		totalDurationMilliseconds += cycle.Inhale.Milliseconds() +
			cycle.InhaleHold.Milliseconds() +
			cycle.Exhale.Milliseconds() +
			cycle.ExhaleHold.Milliseconds()
	}

	return time.Duration(totalDurationMilliseconds) * time.Millisecond
}

// Run the breath cycles
func runBreatheCycles(cycles []BreatheCycle, gaugeChart *widgets.Gauge, renderText func(currentCycleCount int), playSound func(soundName string)) {
	for i, cycle := range cycles {
		renderText(i + 1)
		runBreatheSubCycle("Inhale", cycle.Inhale, gaugeChart, playSound)
		if cycle.InhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.InhaleHold, gaugeChart, playSound)
		}

		runBreatheSubCycle("Exhale", cycle.Exhale, gaugeChart, playSound)
		if cycle.ExhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.ExhaleHold, gaugeChart, playSound)
		}
	}
}

// Run a single breath sub cycle like an inhale or an exhale step
// by waiting the appropriate time and printing information about
// how long is still to go.
func runBreatheSubCycle(subCycleWord string, duration time.Duration, gaugeChart *widgets.Gauge, playSound func(soundName string)) {
	gaugeChart.Label = fmt.Sprintf("%s for %.1f seconds", subCycleWord, float64(duration.Milliseconds())/1000)
	playSound(subCycleWord)
	switch subCycleWord {
	case "Inhale":
		gaugeChart.BarColor = ui.ColorGreen
	case "Exhale":
		gaugeChart.BarColor = ui.ColorBlue
	case "Hold":
		gaugeChart.BarColor = ui.ColorYellow
	}
	gaugeChart.Percent = 0
	ui.Render(gaugeChart)

	for i := int(duration.Milliseconds() / 100); i > 0; i-- {
		time.Sleep(100 * time.Millisecond)

		// Don't play the sound for the first second of a new step
		// because the step word is played then
		firstSecond := int(duration.Milliseconds()/100) - 10 - int(duration.Milliseconds()/100)%10 + 1
		if i < firstSecond && i%10 == 0 {
			playSound(fmt.Sprintf("%d", i/10))
		}

		percentage := int(100 - float64(i)/float64(duration.Milliseconds()/100)*100)
		gaugeChart.Percent = percentage
		ui.Render(gaugeChart)
	}
}

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
