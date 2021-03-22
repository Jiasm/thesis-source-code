package entity

type TaskParticipant struct {
	TaskId  uint `json:"task_id"`
	Uid     uint `json:"uid"`
	AddDate uint `json:"add_date"`
}
