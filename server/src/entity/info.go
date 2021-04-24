package entity

type NewTask struct {
	NewCount    uint   	`json:"new_count"`
	Date     	string	`json:"date"`
}

type TypedTask struct {
	TypeCount   uint   	`json:"type_count"`
	Type     	uint	`json:"type"`
}

type TodoCount struct {
	TodoCount   uint   	`json:"todo_count"`
	Executor    uint	`json:"executor"`
	Status		uint	`json:"status"`
}

type DoneCount struct {
	DoneCount   uint   	`json:"done_count"`
	Executor    uint	`json:"executor"`
}
