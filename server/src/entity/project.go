package entity

type Project struct {
	ID uint `json:"id"`
	GroupId uint `json:"group_id"`
	Creator uint `json:"creator"`
	Name uint `json:"name"`
	Status uint `json:"status"`
}
