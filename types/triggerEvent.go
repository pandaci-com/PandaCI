package types

type TriggerEventSource string

type Committer struct {
	Email  *string
	UserID *string
}

type TriggerEvent struct {
	Source           TriggerEventSource
	SHA              string
	GitTitle         string
	Branch           string
	TargetBranch     *string
	PrNumber         *int32
	Trigger          RunTrigger
	Committer        Committer
	RequiresApproval bool
}

const (
	TriggerEventSourceGithub = "github"
	TriggerEventSourceManual = "manual"
)
