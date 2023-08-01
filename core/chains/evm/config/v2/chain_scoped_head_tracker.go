package v2

import "time"

type headTrackerConfig struct {
	c HeadTracker
}

func (h *headTrackerConfig) HistoryDepth() uint32 {
	return *h.c.HistoryDepth
}

func (h *headTrackerConfig) MaxBufferSize() uint32 {
	return *h.c.MaxBufferSize
}

func (h *headTrackerConfig) SamplingInterval() time.Duration {
	return h.c.SamplingInterval.Duration()
}

func (h *headTrackerConfig) TotalHeadsLimit() int {
	return *h.c.TotalHeadsLimit
}

func (h *headTrackerConfig) PersistHeads() bool {
	return *h.c.PersistHeads
}
func (h *headTrackerConfig) CanonicalChainRule() string {
	return *h.c.CanonicalChainRule
}