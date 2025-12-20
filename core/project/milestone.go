package project

type MilestoneStatus string

const (
	MilestonePending   MilestoneStatus = "pending"
	MilestoneActive    MilestoneStatus = "active"
	MilestoneCompleted MilestoneStatus = "completed"
	MilestoneRejected  MilestoneStatus = "rejected"
)

type Milestone struct {
	ID          string
	Title       string
	Description string
	Budget      int64
	Status      MilestoneStatus
	Executor    string
	Proof       []string
}
