package entity

type GroupMember struct {
	GroupId uint `json:"group_id"`
	Uid uint `json:"uid"`
	RoleId uint `json:"role_id"`
	Status uint `json:"status"`
}
