package breathe

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// A BreathCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreathCycle struct {
	Inhale     time.Duration
	InhaleHold time.Duration
	Exhale     time.Duration
	ExhaleHold time.Duration
}

// Take one BreathCycle and create a slice of the same BreathCycles with
// the cyclesCount of it.
func GenerateBreathCycles(cycle BreathCycle, cyclesCount int) (cycles []BreathCycle) {
	for i := 0; i < cyclesCount; i++ {
		cycles = append(cycles, cycle)
	}
	return cycles
}

// Draw the UI and kick off the breath cycles
func RunBreathCycles(title string, cycles []BreathCycle, sound string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize TermUI: %v", err)
	}
	defer ui.Close()

	gaugeChart := initGaugeChart()
	textBox := initTextBox()
	renderText := createRenderText(textBox, title, cycles)
	playSound := initSpeaker(sound)

	go runBreathCycles(cycles, gaugeChart, renderText, playSound)

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
func createRenderText(textBox *widgets.Paragraph, title string, cycles []BreathCycle) func(currentCycleCount int) {
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
func totalDuration(cycles []BreathCycle) time.Duration {
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
func runBreathCycles(cycles []BreathCycle, gaugeChart *widgets.Gauge, renderText func(currentCycleCount int), playSound func(soundName string)) {
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
		gaugeChart.Percent = 0
	case "Exhale":
		gaugeChart.BarColor = ui.ColorBlue
		gaugeChart.Percent = 100
	case "Hold":
		gaugeChart.BarColor = ui.ColorYellow
		gaugeChart.Percent = 0
	}
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
		if subCycleWord == "Exhale" {
			percentage = 100 - percentage
		}
		gaugeChart.Percent = percentage
		ui.Render(gaugeChart)
	}
}
