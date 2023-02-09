package cmd

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

type Spinner struct {
	quit  chan bool
	apiar *progressbar.ProgressBar
}

func NewSpinner() *Spinner {
	return &Spinner{
		quit: make(chan bool),
		apiar: progressbar.NewOptions(-1,
			progressbar.OptionSpinnerType(14),
			progressbar.OptionClearOnFinish(),
			progressbar.OptionFullWidth(),
		),
	}
}

func (s *Spinner) Start() {
	go func() {
		for {
			select {
			case <-s.quit:
				return
			default:
				s.apiar.Add(1)
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.quit <- true
}
