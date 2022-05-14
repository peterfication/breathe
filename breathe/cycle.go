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
func GenerateBreathCycles(cycle BreathCycle, cyclesCount int) (breathCycles []BreathCycle) {
	for i := 0; i < cyclesCount; i++ {
		breathCycles = append(breathCycles, cycle)
	}
	return breathCycles
}

type Runner struct {
	title string
	sound string

	breathCycles      []BreathCycle
	currentCycleCount int

	// Widgets
	gaugeChart *widgets.Gauge
	textBox    *widgets.Paragraph

	// Closures
	playSound func(soundName string)
}

func (runner *Runner) Init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize TermUI: %v", err)
	}

	runner.InitGaugeChart()
	runner.InitTextBox()
	runner.InitSpeaker(runner.sound)

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
	runner := Runner{title: title, breathCycles: breathCycles, sound: sound}
	runner.Init()
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
	for _, cycle := range runner.breathCycles {
		totalDurationMilliseconds += cycle.Inhale.Milliseconds() +
			cycle.InhaleHold.Milliseconds() +
			cycle.Exhale.Milliseconds() +
			cycle.ExhaleHold.Milliseconds()
	}

	return time.Duration(totalDurationMilliseconds) * time.Millisecond
}

// The total count of breath cycles
func (runner *Runner) BreathCyclesCount() int {
	return len(runner.breathCycles)
}

// Run the breath cycles
func (runner *Runner) RunBreathCycles() {
	for i, cycle := range runner.breathCycles {
		runner.currentCycleCount = i + 1
		runner.RefreshText()
		runBreatheSubCycle("Inhale", cycle.Inhale, runner.gaugeChart, runner.playSound)
		if cycle.InhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.InhaleHold, runner.gaugeChart, runner.playSound)
		}

		runBreatheSubCycle("Exhale", cycle.Exhale, runner.gaugeChart, runner.playSound)
		if cycle.ExhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.ExhaleHold, runner.gaugeChart, runner.playSound)
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
