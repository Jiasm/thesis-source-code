package entity

type Task struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Status uint `json:"status"`
	Creator uint `json:"creator"`
	ExpireDate uint `json:"expire_date"`
	TaskGroupId uint `json:"task_group_id"`
	ParentTaskId uint `json:"parent_task_id"`
	Type uint `json:"type"`
	Priority uint `json:"priority"`
}
