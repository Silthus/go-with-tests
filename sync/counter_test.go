package sync

import (
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type CounterTestSuite struct {
	suite.Suite
	Counter *Counter
}

func (s *CounterTestSuite) SetupTest() {
	s.Counter = &Counter{}
}

func (s *CounterTestSuite) TestCounter() {
	s.Run("incrementing counter 3 times leaves it at 3", func() {
		s.increaseCount(3)
		s.assertCounter(3)
	})
}

func (s *CounterTestSuite) TestCounterConcurrency() {
	s.Run("it runs safely concurrently", func() {
		wantedCount := 1000
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				s.Counter.Increase()
				wg.Done()
			}()
		}
		wg.Wait()

		s.assertCounter(wantedCount)
	})
}

func TestCounterTestSuite(t *testing.T) {
	suite.Run(t, new(CounterTestSuite))
}

func (s *CounterTestSuite) assertCounter(count int) {
	s.T().Helper()
	s.Equal(count, s.Counter.Value())
}

func (s *CounterTestSuite) increaseCount(times int) {
	for i := 0; i < times; i++ {
		s.Counter.Increase()
	}
}
