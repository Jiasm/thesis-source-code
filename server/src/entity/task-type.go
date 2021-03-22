package entity

type TaskType struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
	Desc string `json:"desc"`
}
