package protocol

type ProposalSink interface {
	AddProposal(intent ProposalIntent) error
}
