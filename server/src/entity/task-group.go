package entity

type TaskGroup struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Status uint `json:"status"`
	Creator uint `json:"creator"`
}
