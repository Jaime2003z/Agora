package node

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// handleStatus returns the status of the node
func (n *Node) handleStatus(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "ok",
		"node_id":   n.ID(),
		"protocols": n.Host.Mux().Protocols(),
	})
}

// handleListProposals returns a list of all proposals
func (n *Node) handleListProposals(w http.ResponseWriter, r *http.Request) {
	proposals := n.Governance.ListProposals()
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"proposals": proposals,
		"count":     len(proposals),
	})
}

// handleGetProposal returns a specific proposal by ID
func (n *Node) handleGetProposal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	proposalID := vars["id"]

	proposal, exists := n.Governance.GetProposal(proposalID)
	if !exists {
		respondError(w, http.StatusNotFound, "Proposal not found")
		return
	}

	votes, _ := n.Governance.GetVotes(proposalID)

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"proposal": proposal,
		"votes":    votes,
	})
}

// handleCloseProposal closes a proposal and calculates the result
func (n *Node) handleCloseProposal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	proposalID := vars["id"]

	result, err := n.Governance.CloseProposal(proposalID)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "proposal not found" {
			status = http.StatusNotFound
		} else if err.Error() == "proposal is already closed" {
			status = http.StatusBadRequest
		}
		respondError(w, status, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Proposal closed successfully",
		"result":  result,
	})
}

// handleCreateProposal handles proposal creation
func (n *Node) handleCreateProposal(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CreatorID   string `json:"creator_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	proposal, err := n.Governance.CreateProposal(req.Title, req.Description, req.CreatorID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"message":  "Proposal created successfully",
		"proposal": proposal,
	})
}

// handleVote handles voting on a proposal
func (n *Node) handleVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	proposalID := vars["id"]

	var req struct {
		VoterID  string `json:"voter_id"`
		Approved bool   `json:"approved"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	vote, err := n.Governance.AddVote(proposalID, req.VoterID, req.Approved)
	if err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "proposal not found":
			status = http.StatusNotFound
		case "cannot vote on closed proposal", "voter has already voted on this proposal":
			status = http.StatusBadRequest
		}
		respondError(w, status, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Vote cast successfully",
		"vote":    vote,
	})
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError sends a JSON error response
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
