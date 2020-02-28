package systemtray

import (
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/meinto/cntr/counter"
)

type Systemtray struct {
	counter *counter.Counter
}

func NewSystemtrayWidget(c *counter.Counter) *Systemtray {
	return &Systemtray{c}
}

func (s *Systemtray) Run() {
	systray.Run(s.onReady, s.onExit)
}

func (s *Systemtray) onReady() {
	systray.SetTitle("Key count: 0")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	for {
		systray.SetTitle("Key count: " + strconv.FormatInt(int64(s.counter.GetKeys()), 10))
		time.Sleep(time.Second)
	}
}

func (s *Systemtray) onExit() {
	// clean up here
}