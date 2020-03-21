package main

import (
	"sync"

	"github.com/nitishm/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	// New instance of the logger
	l := logger.New()
	l.AddField("main", 0)
	l.SetDefaults(logrus.Fields{"default_0": "def_main"})

	wg := &sync.WaitGroup{}
	// Start first go-routine that uses the global logger
	// This should result in unpredictable output from
	// the main routine and the go-routine since they
	// work on the same logger instance and depends on
	// who get the mutex lock first
	wg.Add(2)
	go func() {
		l.Info("Starting first goroutine")
		l.AddField("first", 1)
		l.Info("Log to console from first go routine")
		wg.Done()
	}()

	// Start second go-routine with a clone or shallow
	// copy of the logger. Clone() creates a new instance
	// of the logger but copies over the fields and defaults
	// from the parent
	go func(ls logger.Logger) {
		ls.Info("Starting second goroutine")
		ls.SetDefaults(logrus.Fields{"default_2": "def_second"})
		ls.AddField("second", 2)
		ls.Info("Log to console from second go routine")
		wg.Done()
	}(l.Clone())

	l.Info("Log to console from main")
	l.ResetFields()
	l.Info("Log to console from main after reset")

	wg.Wait()
}
