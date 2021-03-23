package entity

type Project struct {
	ID      uint 	`json:"id"`
	GroupId uint 	`json:"group_id"`
	Creator uint 	`json:"creator"`
	Name    string  `json:"name"`
	Status  uint 	`json:"status"`
}
