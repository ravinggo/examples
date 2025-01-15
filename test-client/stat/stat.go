package stat

import (
	"sync/atomic"
	"time"
)

type Stat struct {
	Total         int64
	TotalDuration int64
}

func (s *Stat) Add(d time.Duration) {
	atomic.AddInt64(&s.Total, 1)
	atomic.AddInt64(&s.TotalDuration, d.Nanoseconds())

}

func (s *Stat) Avg() time.Duration {
	if s.Total == 0 {
		return 0
	}
	return time.Duration(s.TotalDuration / s.Total)
}

type Stats struct {
	Stats map[int]*Stat
}

func NewStats() *Stats {
	return &Stats{
		Stats: make(map[int]*Stat),
	}
}

func (s *Stats) Add(name int, d time.Duration) {
	if _, ok := s.Stats[name]; !ok {
		s.Stats[name] = &Stat{}
	}
	s.Stats[name].Add(d)
}

func (s *Stats) Avg() time.Duration {
	var total time.Duration
	for _, stat := range s.Stats {
		total += stat.Avg()
	}
	return total / time.Duration(len(s.Stats))
}
