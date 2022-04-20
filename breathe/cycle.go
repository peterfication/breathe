package breathe

import (
	"fmt"
	"log"
	"time"

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
func RunBreatheCycles(title string, cycles []BreatheCycle) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize TermUI: %v", err)
	}
	defer ui.Close()

	termWidth, termHeight := ui.TerminalDimensions()

	gaugeChart := widgets.NewGauge()
	gaugeChart.Title = fmt.Sprintf("%s (total duration: %s)", title, TotalDuration(cycles))
	gaugeChart.SetRect(0, 0, termWidth, termHeight)
	gaugeChart.Percent = 0
	gaugeChart.BarColor = ui.ColorGreen
	gaugeChart.BorderStyle.Fg = ui.ColorWhite
	gaugeChart.TitleStyle.Fg = ui.ColorCyan
	gaugeChart.Label = "Inhale"

	ui.Render(gaugeChart)

	go runBreatheCycles(cycles, gaugeChart)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

// Run the breath cycles
func runBreatheCycles(cycles []BreatheCycle, gaugeChart *widgets.Gauge) {
	for _, cycle := range cycles {
		runBreatheSubCycle("Inhale", cycle.Inhale, gaugeChart)
		if cycle.InhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.InhaleHold, gaugeChart)
		}

		runBreatheSubCycle("Exhale", cycle.Exhale, gaugeChart)
		if cycle.ExhaleHold.Milliseconds() > 0 {
			runBreatheSubCycle("Hold", cycle.ExhaleHold, gaugeChart)
		}
	}
}

// Run a single breath sub cycle like an inhale or an exhale step
// by waiting the appropriate time and printing information about
// how long is still to go.
func runBreatheSubCycle(subCycleWord string, duration time.Duration, gaugeChart *widgets.Gauge) {
	gaugeChart.Label = fmt.Sprintf("%s for %.1f seconds", subCycleWord, float64(duration.Milliseconds())/1000)
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

		percentage := int(100 - float64(i)/float64(duration.Milliseconds()/100)*100)
		gaugeChart.Percent = percentage
		ui.Render(gaugeChart)
	}
}

func TotalDuration(cycles []BreatheCycle) time.Duration {
	totalDurationMilliseconds := int64(0)
	for _, cycle := range cycles {
		totalDurationMilliseconds += cycle.Inhale.Milliseconds() +
			cycle.InhaleHold.Milliseconds() +
			cycle.Exhale.Milliseconds() +
			cycle.ExhaleHold.Milliseconds()
	}

	return time.Duration(totalDurationMilliseconds) * time.Millisecond
}
