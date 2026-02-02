package main

import (
	"github.com/jhlee-young/delapore-mouse/internal/app"
	"github.com/jhlee-young/delapore-mouse/internal/config"
)

func main() {
	cfg := config.DefaultConfig()
	w := app.NewWindow(cfg)
	w.ShowAndRun()
}
