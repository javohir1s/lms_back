package models

type Schedule struct {
	Id         string `json:"id"`
	Group_id   string `json:"group_id"`
	Group_type string `json:"group_type"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	Date       string `json:"date"`
	Branch_id  string `json:"branch_id"`
	Teacher_id string `json:"teacher_id"`
	Created_at string `json:"created_id"`
	Updated_at string `json:"updated_id"`
}

type GetAllSchedulesResponse struct {
	Schedules []Schedule `json:"students"`
	Count    int16     `json:"count"`
}

type GetAllSchedulesRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
