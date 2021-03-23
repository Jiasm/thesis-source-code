package entity

type Group struct {
	ID         	uint   	`json:"id"`
	Name     	string	`json:"name"`
	Status     	uint   	`json:"status"`
	Creator    	uint	`json:"creator"`
	CreatedDate uint 	`json:"created_date""`
}
