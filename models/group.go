package models

type Group struct {
	Id         string `json:"id"`
	Group_id   string `json:"group_id"`
	Branch_id  string `json:"branch_id"`
	Teacher_id string `json:"teacher_id"`
	Type       string `json:"type"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type GetAllGroupsResponse struct {
	Groups []Group `json:"groups"`
	Count    int16     `json:"count"`
}

type GetAllGroupsRequest struct {
    Search string `json:"search"`
	Page uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}