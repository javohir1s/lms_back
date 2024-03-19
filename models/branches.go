package models

type Branches struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type GetAllBranchesResponse struct {
	Branches []Branches `json:"branches"`
	Count    int16      `json:"count"`
}

type GetAllBranchesRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
