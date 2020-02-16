package timer

import (
	"github.com/tomoyane/grant-n-z/gnzcache/service"
	"time"

	"github.com/tomoyane/grant-n-z/gnz/log"
)

// UpdateTimer interface
type UpdateTimer interface {
	// Start update cache timer
	Start(exitCode chan int) int

	// Stop update cache timer
	Stop()
}

// UpdateTimer struct
type UpdateTimerImpl struct {
	Ticker         *time.Ticker
	IsRunning      bool
	UpdaterService service.UpdaterService
}

// Constructor
func NewUpdateTimer() UpdateTimer {
	return UpdateTimerImpl{
		Ticker:         time.NewTicker(5 * time.Minute),
		IsRunning:      true,
		UpdaterService: service.NewUpdaterService(),
	}
}

// Start cache update timer
func (ut UpdateTimerImpl) Start(exitCode chan int) int {
	code := 0

loop:
	for {
		select {
		case <-ut.Ticker.C:
			// TODO
		case c := <-exitCode:
			log.Logger.Info("Break update cache loop")
			ut.Ticker.Stop()
			code = c
			break loop
		}
	}

	log.Logger.Info("Stopped update cache process")
	return code
}

func (ut UpdateTimerImpl) Stop() {
	ut.Ticker.Stop()
}