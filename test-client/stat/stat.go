package stat

import (
	"sync/atomic"
	"time"
)

type Stat struct {
	Total         int64
	TotalDuration int64
}

func NewStat() *Stat {
	return &Stat{}
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
