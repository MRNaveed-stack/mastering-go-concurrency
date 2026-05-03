package counter

import "sync"

type ResultStats struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewStats() *ResultStats {
	return &ResultStats{
		counts: make(map[string]int),
	}
}

func (s *ResultStats) Update(word string, count int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counts[word] += count
}

// Counts returns a copy of the current word counts.
func (s *ResultStats) Counts() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]int, len(s.counts))
	for k, v := range s.counts {
		out[k] = v
	}
	return out
}
