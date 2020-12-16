package stopper

import (
	"sync"
)

type Stopper struct {
	wg sync.WaitGroup
}

func (s *Stopper) AddWorker(f func()) {
	s.wg.Add(1)
	go func() {
		f()
		s.wg.Done()
	}()
}
