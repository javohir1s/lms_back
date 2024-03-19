package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"lms_back/models"
	"lms_back/pkg"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranch(db *sql.DB) branchRepo {
	return branchRepo{
		db: db,
	}
}

func (c *branchRepo) Create(branch models.Branches) (string, error) {

	id := uuid.New()
	query := `INSERT INTO branches (
		id,
		name,
		address,
		created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		branch.Name,
		branch.Address)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c *branchRepo) Update(branch models.Branches) (string, error) {
	query := `update branches set 
	name=$1,
	address=$2,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $3 AND deleted_at = 0
	`
	_, err := c.db.Exec(query,
		branch.Name,
		branch.Address,
		branch.Id,
	)
	if err != nil {
		return "", err
	}
	return branch.Id, nil
}

func (c *branchRepo) GetAll(req models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {
	var (
		resp   = models.GetAllBranchesResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select count(id) over(),
        id,
        name,
        address,
        created_at,
        updated_at,
        deleted_at FROM branches WHERE deleted_at = 0` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			branch   = models.Branches{}
			updateAt sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.CreatedAt,
			&updateAt,
			&branch.DeletedAt); err != nil {
			return resp, err
		}
		branch.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Branches = append(resp.Branches, branch)
	}
	return resp, nil
}

func (c *branchRepo) GetByID(id string) (models.Branches, error) {
	branch := models.Branches{}

	if err := c.db.QueryRow(`select id, name, address, created_at from branches where id = $1`, id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt); err != nil {
		return models.Branches{}, err
	}
	return branch, nil
}

func (c *branchRepo) Delete(id string) error {
	query := `delete from branches where id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
