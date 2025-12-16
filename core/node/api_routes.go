package node

import "github.com/gorilla/mux"

// SetupAPIRoutes configures all API routes
func (n *Node) SetupAPIRoutes(router *mux.Router) {
	// Status endpoint
	router.HandleFunc("/status", n.handleStatus).Methods("GET")

	// Governance endpoints
	router.HandleFunc("/governance/proposals", n.handleCreateProposal).Methods("POST")
	router.HandleFunc("/governance/proposals", n.handleListProposals).Methods("GET")
	router.HandleFunc("/governance/proposals/{id}", n.handleGetProposal).Methods("GET")
	router.HandleFunc("/governance/proposals/{id}/votes", n.handleVote).Methods("POST")
	router.HandleFunc("/governance/proposals/{id}/close", n.handleCloseProposal).Methods("POST")
}
