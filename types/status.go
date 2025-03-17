package types

import (
	"fmt"

	pb "github.com/pandaci-com/pandaci/proto/go/v1"
)

type RunConclusion string

const (
	RunConclusionSuccess   RunConclusion = "success"
	RunConclusionFailure   RunConclusion = "failure"
	RunConclusionSkipped   RunConclusion = "skipped"
	RunConclusionCancelled RunConclusion = "cancelled"
)

var WorstRunConclusion = RunConclusion(RunConclusionFailure)

var RunConclusionRank = map[RunConclusion]int{
	RunConclusionFailure: 3,
	RunConclusionSuccess: 2,
	RunConclusionSkipped: 1,
}

var RunConclusionRankProto = map[pb.Conclusion]int{
	pb.Conclusion_CONCLUSION_UNSPECIFIED: 0,
	pb.Conclusion_CONCLUSION_SKIPPED:     1,
	pb.Conclusion_CONCLUSION_SUCCESS:     2,
	pb.Conclusion_CONCLUSION_FAILURE:     3,
}

func CompareProtoConclusionRank(conclusion1, conclusion2 pb.Conclusion) bool {
	return RunConclusionRankProto[conclusion1] > RunConclusionRankProto[conclusion2]
}

func RunOutputFromProto(conclusion pb.Conclusion) (RunConclusion, error) {

	switch conclusion {
	case pb.Conclusion_CONCLUSION_FAILURE:
		return RunConclusionFailure, nil
	case pb.Conclusion_CONCLUSION_SUCCESS:
		return RunConclusionSuccess, nil
	case pb.Conclusion_CONCLUSION_SKIPPED:
		return RunConclusionSkipped, nil
	default:
		return "", fmt.Errorf("Unable to convert output proto %v to types.RunConclusion", conclusion)
	}

}

type RunStatus string

const (
	RunStatusQueued           RunStatus = "queued"
	RunStatusRunning          RunStatus = "running"
	RunStatusCompleted        RunStatus = "completed"
	RunStatusPending          RunStatus = "pending"
	RunStatusAwaitingApproval RunStatus = "awaiting-approval"
)
