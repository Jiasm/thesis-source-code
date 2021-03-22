package entity

type ProjectMember struct {
	ProjectId uint `json:"project_id"`
	Uid       uint `json:"uid"`
	RoleId    uint `json:"role_id"`
	Status    uint `json:"status"`
}
