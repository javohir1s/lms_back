package models

type Student struct {
	ID         string  `json:"id"`
	Full_Name  string  `json:"full_name"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	PaidSum    float64 `json:"paid_sum"`
	Status     string  `json:"status"`
	Login      string  `json:"login"`
	Password   string  `json:"password"`
	GroupID    string  `json:"group_id"`
	Created_At string  `json:"created_at"`
	Updated_At string  `json:"updated_at"`
	Deleted_At int     `json:"deleted_at"`
}

type GetAllStudentsResponse struct {
	Students []Student `json:"students"`
	Count    int16     `json:"count"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
