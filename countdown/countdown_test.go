package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestPrintCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})
		assert.Equal(t, `3
2
1
Go!`, buffer.String())
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		assert.Equal(t, []string{
			sleep, write,
			sleep, write,
			sleep, write,
			sleep, write,
		}, spySleepPrinter.Calls)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	spyTime := &SpyTime{}
	duration := 5 * time.Second
	sleeper := ConfigurableSleeper{
		duration: duration,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()
	assert.Equal(t, duration, spyTime.durationSlept)
}
