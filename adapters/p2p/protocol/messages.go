package protocol

type MessageType string

const (
	MessageProjectProposal MessageType = "PROJECT_PROPOSAL"
	MessageVote            MessageType = "VOTE"
	MessageSyncRequest     MessageType = "SYNC_REQUEST"
	MessageSyncResponse    MessageType = "SYNC_RESPONSE"
)

type Message struct {
	ID        string
	Type      MessageType
	Payload   []byte
	Timestamp int64
	SenderID  string
}
