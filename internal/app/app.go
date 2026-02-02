package app

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jhlee-young/delapore-mouse/internal/config"
)

var (
	running bool
	mu      sync.Mutex
)

func NewWindow(cfg config.Config) fyne.Window {
	a := app.New()
	w := a.NewWindow("Delapore Mouse")
	w.Resize(fyne.NewSize(250, 120))

	startBtn := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), nil)
	stopBtn := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), nil)
	stopBtn.Disable()

	startBtn.OnTapped = func() {
		mu.Lock()
		if running {
			mu.Unlock()
			return
		}
		running = true
		mu.Unlock()

		startBtn.Disable()
		stopBtn.Enable()
	}

	stopBtn.OnTapped = func() {
		mu.Lock()
		running = false
		mu.Unlock()

		stopBtn.Disable()
		startBtn.Enable()
	}

	buttons := container.NewHBox(startBtn, stopBtn)
	centered := container.NewCenter(buttons)
	padded := container.New(layout.NewPaddedLayout(), centered)

	w.SetContent(padded)
	return w
}
