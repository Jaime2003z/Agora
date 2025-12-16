// core/governance/proposal.go
package governance

import (
	"database/sql"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Jaime2003z/Agora/core/storage"
)

type Proposal struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	CreatorID   string          `json:"creator_id"`
	CreatedAt   int64           `json:"created_at"`
	Closed      bool            `json:"closed"`
	Result      *ProposalResult `json:"result,omitempty"`
}

type ProposalResult struct {
	Approved   bool  `json:"approved"`
	YesVotes   int   `json:"yes_votes"`
	NoVotes    int   `json:"no_votes"`
	TotalVotes int   `json:"total_votes"`
	DecidedAt  int64 `json:"decided_at"`
}

type GovernanceService struct {
	db *storage.DB
}

func NewGovernanceService(db *storage.DB) *GovernanceService {
	return &GovernanceService{db: db}
}

func (s *GovernanceService) CreateProposal(title, description, creatorID string) (*Proposal, error) {
	proposal := &Proposal{
		ID:          generateID("prop"),
		Title:       title,
		Description: description,
		CreatorID:   creatorID,
		CreatedAt:   time.Now().Unix(),
		Closed:      false,
	}

	_, err := s.db.Exec(`
		INSERT INTO proposals (id, title, description, creator_id, created_at, closed)
		VALUES (?, ?, ?, ?, ?, ?)
	`, proposal.ID, proposal.Title, proposal.Description, proposal.CreatorID, proposal.CreatedAt, false)

	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (s *GovernanceService) GetProposal(id string) (*Proposal, bool) {
	var p Proposal
	var result ProposalResult
	var resultApproved sql.NullBool
	var resultYes, resultNo, resultTotal sql.NullInt64
	var decidedAt sql.NullInt64

	err := s.db.QueryRow(`
		SELECT id, title, description, creator_id, created_at, closed,
		       result_approved, result_yes_votes, result_no_votes, result_total_votes, result_decided_at
		FROM proposals WHERE id = ?
	`, id).Scan(
		&p.ID, &p.Title, &p.Description, &p.CreatorID, &p.CreatedAt, &p.Closed,
		&resultApproved, &resultYes, &resultNo, &resultTotal, &decidedAt,
	)

	if err == sql.ErrNoRows {
		return nil, false
	}
	if err != nil {
		return nil, false
	}

	if resultApproved.Valid {
		result = ProposalResult{
			Approved:   resultApproved.Bool,
			YesVotes:   int(resultYes.Int64),
			NoVotes:    int(resultNo.Int64),
			TotalVotes: int(resultTotal.Int64),
			DecidedAt:  decidedAt.Int64,
		}
		p.Result = &result
	}

	return &p, true
}

func (s *GovernanceService) AddVote(proposalID, voterID string, approved bool) (*Vote, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check if proposal exists and is not closed
	var closed bool
	err = tx.QueryRow("SELECT closed FROM proposals WHERE id = ?", proposalID).Scan(&closed)
	if err == sql.ErrNoRows {
		return nil, errors.New("proposal not found")
	}
	if err != nil {
		return nil, err
	}
	if closed {
		return nil, errors.New("cannot vote on closed proposal")
	}

	// Check if voter has already voted
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM votes WHERE proposal_id = ? AND voter_id = ?", proposalID, voterID).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("voter has already voted on this proposal")
	}

	// Add vote
	vote := Vote{
		ID:         generateID("vote"),
		ProposalID: proposalID,
		VoterID:    voterID,
		Approved:   approved,
		Timestamp:  time.Now().Unix(),
	}

	_, err = tx.Exec(`
		INSERT INTO votes (id, proposal_id, voter_id, approved, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`, vote.ID, vote.ProposalID, vote.VoterID, vote.Approved, vote.Timestamp)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &vote, nil
}

func (s *GovernanceService) CloseProposal(proposalID string) (*ProposalResult, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Get proposal
	var p Proposal
	var closed bool
	err = tx.QueryRow(`
		SELECT id, title, description, creator_id, created_at, closed 
		FROM proposals WHERE id = ?`, proposalID).Scan(
		&p.ID, &p.Title, &p.Description, &p.CreatorID, &p.CreatedAt, &closed,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("proposal not found")
	}
	if err != nil {
		return nil, err
	}
	if closed {
		return nil, errors.New("proposal is already closed")
	}

	// Count votes
	var result struct {
		Yes int
		No  int
	}
	rows, err := tx.Query(`
		SELECT COUNT(*) as count, approved 
		FROM votes 
		WHERE proposal_id = ? 
		GROUP BY approved`, proposalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		var approved bool
		if err := rows.Scan(&count, &approved); err != nil {
			return nil, err
		}
		if approved {
			result.Yes = count
		} else {
			result.No = count
		}
	}

	totalVotes := result.Yes + result.No
	resultObj := &ProposalResult{
		Approved:   result.Yes > result.No,
		YesVotes:   result.Yes,
		NoVotes:    result.No,
		TotalVotes: totalVotes,
		DecidedAt:  time.Now().Unix(),
	}

	// Update proposal
	_, err = tx.Exec(`
		UPDATE proposals 
		SET closed = 1, 
		    result_approved = ?,
		    result_yes_votes = ?,
		    result_no_votes = ?,
		    result_total_votes = ?,
		    result_decided_at = ?
		WHERE id = ?`,
		resultObj.Approved,
		resultObj.YesVotes,
		resultObj.NoVotes,
		resultObj.TotalVotes,
		resultObj.DecidedAt,
		proposalID,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return resultObj, nil
}

func (s *GovernanceService) GetVotes(proposalID string) ([]Vote, bool) {
	rows, err := s.db.Query(`
		SELECT id, proposal_id, voter_id, approved, timestamp
		FROM votes
		WHERE proposal_id = ?
		ORDER BY timestamp`, proposalID)
	if err != nil {
		return nil, false
	}
	defer rows.Close()

	var votes []Vote
	for rows.Next() {
		var v Vote
		if err := rows.Scan(&v.ID, &v.ProposalID, &v.VoterID, &v.Approved, &v.Timestamp); err != nil {
			return nil, false
		}
		votes = append(votes, v)
	}

	return votes, true
}

func (s *GovernanceService) ListProposals() []*Proposal {
	rows, err := s.db.Query(`
		SELECT id, title, description, creator_id, created_at, closed,
		       result_approved, result_yes_votes, result_no_votes, result_total_votes, result_decided_at
		FROM proposals
		ORDER BY created_at DESC`)
	if err != nil {
		return []*Proposal{}
	}
	defer rows.Close()

	var proposals []*Proposal
	for rows.Next() {
		var p Proposal
		var resultApproved sql.NullBool
		var resultYes, resultNo, resultTotal sql.NullInt64
		var decidedAt sql.NullInt64

		err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.CreatorID, &p.CreatedAt, &p.Closed,
			&resultApproved, &resultYes, &resultNo, &resultTotal, &decidedAt,
		)
		if err != nil {
			continue
		}

		if resultApproved.Valid {
			p.Result = &ProposalResult{
				Approved:   resultApproved.Bool,
				YesVotes:   int(resultYes.Int64),
				NoVotes:    int(resultNo.Int64),
				TotalVotes: int(resultTotal.Int64),
				DecidedAt:  decidedAt.Int64,
			}
		}

		proposals = append(proposals, &p)
	}

	return proposals
}

var idCounter int64

func generateID(prefix string) string {
	return fmt.Sprintf("%s_%d_%d", prefix, time.Now().Unix(), atomic.AddInt64(&idCounter, 1))
}
