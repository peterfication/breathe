package breathe

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Runner struct {
	title string

	breathCycles      []BreathCycle
	currentCycleCount int

	// Widgets
	gaugeChart *widgets.Gauge
	textBox    *widgets.Paragraph

	speaker Speaker
}

func (runner *Runner) Init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize TermUI: %v", err)
	}

	runner.InitGaugeChart()
	runner.InitTextBox()

	runner.Render()
}

func (runner *Runner) Run() {
	go runner.RunBreathCycles()

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func (runner *Runner) Close() {
	ui.Close()
}

// Init the Runner, run it and close it afterwards
func RunBreathCycles(title string, breathCycles []BreathCycle, sound string) {
	runner := Runner{title: title, breathCycles: breathCycles, speaker: Speaker{sound: sound}}
	runner.Init()
	runner.speaker = NewSpeaker(sound)
	defer runner.Close()

	runner.Run()
}

// The gauge chart that displays the breathing action to perform
func (runner *Runner) InitGaugeChart() {
	termWidth, termHeight := ui.TerminalDimensions()
	gaugeChart := widgets.NewGauge()
	gaugeChart.SetRect(0, 0, termWidth, termHeight-10)
	gaugeChart.Percent = 0
	gaugeChart.BarColor = ui.ColorGreen
	gaugeChart.BorderStyle.Fg = ui.ColorWhite
	gaugeChart.TitleStyle.Fg = ui.ColorCyan

	runner.gaugeChart = gaugeChart
}

// The text box that holds additional information about the breathing cycles
func (runner *Runner) InitTextBox() {
	termWidth, termHeight := ui.TerminalDimensions()
	paragraph := widgets.NewParagraph()
	paragraph.SetRect(0, termHeight-10, termWidth, termHeight)
	paragraph.Text = "Breathe"

	runner.textBox = paragraph
}

// Rerender the UI elements
func (runner *Runner) Render() {
	ui.Render(runner.gaugeChart)
	ui.Render(runner.textBox)
}

// Update the text in the text box
func (runner *Runner) RefreshText() {
	text := fmt.Sprintf(`Always inhale through the nose!

%s
Total duration: %s
Cycle %d of %d
`, runner.title, runner.TotalDuration(), runner.currentCycleCount, runner.BreathCyclesCount())
	runner.textBox.Text = text
	runner.Render()
}

// Sum up the duration of all steps in all cycles
func (runner *Runner) TotalDuration() time.Duration {
	totalDurationMilliseconds := int64(0)
	for _, breathCycle := range runner.breathCycles {
		totalDurationMilliseconds += breathCycle.Inhale.Milliseconds() +
			breathCycle.InhaleHold.Milliseconds() +
			breathCycle.Exhale.Milliseconds() +
			breathCycle.ExhaleHold.Milliseconds()
	}

	return time.Duration(totalDurationMilliseconds) * time.Millisecond
}

// The total count of breath cycles
func (runner *Runner) BreathCyclesCount() int {
	return len(runner.breathCycles)
}

// Run the breath cycles
func (runner *Runner) RunBreathCycles() {
	for i, breathCycle := range runner.breathCycles {
		runner.currentCycleCount = i + 1
		runner.RefreshText()
		runner.RunBreatheSubCycle("Inhale", breathCycle.Inhale)
		if breathCycle.InhaleHold.Milliseconds() > 0 {
			runner.RunBreatheSubCycle("Hold", breathCycle.InhaleHold)
		}

		runner.RunBreatheSubCycle("Exhale", breathCycle.Exhale)
		if breathCycle.ExhaleHold.Milliseconds() > 0 {
			runner.RunBreatheSubCycle("Hold", breathCycle.ExhaleHold)
		}
	}
}

// Run a single breath sub cycle like an inhale or an exhale step
// by waiting the appropriate time and printing information about
// how long is still to go.
func (runner *Runner) RunBreatheSubCycle(subCycleWord string, duration time.Duration) {
	runner.gaugeChart.Label = fmt.Sprintf("%s for %.1f seconds", subCycleWord, float64(duration.Milliseconds())/1000)
	runner.speaker.PlaySound(subCycleWord)
	switch subCycleWord {
	case "Inhale":
		runner.gaugeChart.BarColor = ui.ColorGreen
		runner.gaugeChart.Percent = 0
	case "Exhale":
		runner.gaugeChart.BarColor = ui.ColorBlue
		runner.gaugeChart.Percent = 100
	case "Hold":
		runner.gaugeChart.BarColor = ui.ColorYellow
		runner.gaugeChart.Percent = 0
	}
	runner.Render()

	for i := int(duration.Milliseconds() / 100); i > 0; i-- {
		time.Sleep(100 * time.Millisecond)

		// Don't play the sound for the first second of a new step
		// because the step word is played then
		firstSecond := int(duration.Milliseconds()/100) - 10 - int(duration.Milliseconds()/100)%10 + 1
		if i < firstSecond && i%10 == 0 {
			runner.speaker.PlaySound(fmt.Sprintf("%d", i/10))
		}

		percentage := int(100 - float64(i)/float64(duration.Milliseconds()/100)*100)
		if subCycleWord == "Exhale" {
			percentage = 100 - percentage
		}
		runner.gaugeChart.Percent = percentage
		runner.Render()
	}
}
