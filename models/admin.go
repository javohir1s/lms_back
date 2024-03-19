package models

type Admin struct {
	Id         string `json:"id"`
	Full_Name  string `json:"full_name"`
	Email      string `json:"email"`
	Age        uint   `json:"age"`
	Status     string `json:"status"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Created_at string `json:"create_at"`
	Updated_at string `json:"updated_at"`
}

type GetAllAdminsResponse struct {
	Admins []Admin `json:"admins"`
	Count  int16   `json:"count"`
}
type GetAllAdminsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
