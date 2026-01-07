package consensus

import (
	"github.com/Jaime2003z/Agora/adapters/p2p/mempool"
)

// ==============================
// Engine
// ==============================

// Tick advances the logical clock by one tick
func (e *Engine) Tick() {
	e.ticker.Advance()
}

type Engine struct {
	mempool *mempool.Mempool
	ticker  *Ticker
	final   *Finality
}

// Constructor
func NewEngine(
	mp *mempool.Mempool,
	ticker *Ticker,
	final *Finality,
) *Engine {
	return &Engine{
		mempool: mp,
		ticker:  ticker,
		final:   final,
	}
}

// ==============================
// Ciclo principal
// ==============================

func (e *Engine) Step() error {
	currentTick := e.ticker.Current()

	// 1️⃣ Obtener proposals pendientes
	proposals := e.mempool.Store().ListProposals()

	for _, intent := range proposals {

		// 2️⃣ Verificar si puede ser finalizado
		if !e.final.CanFinalize(intent, currentTick) {
			continue
		}

		// 3️⃣ Aplicar al estado (esto luego va a state/)
		// applyProposal(intent)

		// 4️⃣ Remover del mempool
		e.mempool.Store().RemoveProposal(intent.MessageID)
	}

	return nil
}
