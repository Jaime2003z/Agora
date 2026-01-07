package consensus

import "sync"

// ==============================
// Ticker lógico
// ==============================

type Ticker struct {
	mu   sync.RWMutex
	tick uint64
}

func NewTicker(start uint64) *Ticker {
	return &Ticker{
		tick: start,
	}
}

func (t *Ticker) Current() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tick
}

// Avanza el tiempo lógico (por bloque, ronda, etc.)
func (t *Ticker) Advance() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tick++
}
