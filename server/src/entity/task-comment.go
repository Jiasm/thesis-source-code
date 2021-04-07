package entity

type TaskComment struct {
	ID       	uint   `json:"id"`
	TaskId   	uint   `json:"task_id"`
	Text     	string `json:"text"`
	Status   	uint   `json:"status"`
	Creator  	uint   `json:"creator"`
	CreatedDate uint   `json:"created_date"`
}
