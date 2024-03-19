package models

type Teacher struct {
	Id         string `json:"id"`
	Full_name  string `json:"full_name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Status     string `json:"status"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}
type GetAllTeachersResponse struct {
	Teachers []Teacher `json:"teachers"`
	Count    int16     `json:"count"`
}

type GetAllTeachersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
