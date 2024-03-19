package models

type Payment struct {
	Id         string  `json:"id"`
	Price      float64 `json:"price"`
	Student_id string  `json:"student_id"`
	Branch_id  string  `json:"branch_id"`
	Admin_id   string  `json:"admin_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type GetAllPaymentsResponse struct {
	Payments []Payment `json:"payments"`
	Count    int16     `json:"count"`
}

type GetAllPaymentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
