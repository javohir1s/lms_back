package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"
	"lms_back/pkg"

	"github.com/google/uuid"
)

type adminRepo struct {
	db *sql.DB
}

func NewAdmin(db *sql.DB) adminRepo {
	return adminRepo{
		db: db,
	}
}

func (c *adminRepo) Create(admin models.Admin) (string, error) {

	id := uuid.New()
	query := `INSERT INTO "admin" (
		id,
		full_name,
		email,
		age,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c *adminRepo) Update(admin models.Admin) (string, error) {
	query := `update "admin" set 
	full_name=$1,
	email=$2,
	age=$3,
	status=$4,
	login=$5,
	password=$6,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $7
	`
	_, err := c.db.Exec(query,
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password,
		admin.Id,
	)
	if err != nil {
		return "", err
	}
	return admin.Id, nil
}

func (c *adminRepo) GetAll(req models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {
	var (
		resp   = models.GetAllAdminsResponse{}
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
        full_name,
        email,
		age,
		status,
		login,
		password,
        created_at,
        updated_at
        FROM "admin"` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			admin    = models.Admin{}
			updateAt sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&admin.Id,
			&admin.Full_Name,
			&admin.Email,
			&admin.Age,
			&admin.Status,
			&admin.Login,
			&admin.Password,
			&admin.Created_at,
			&updateAt); err != nil {
			return resp, err
		}
		admin.Updated_at = pkg.NullStringToString(updateAt)
		resp.Admins = append(resp.Admins, admin)
	}
	return resp, nil
}

func (c *adminRepo) GetByID(id string) (models.Admin, error) {
	admin := models.Admin{}

	if err := c.db.QueryRow(`select id, full_name, email, age, status, login, password, created_at from "admin" where id = $1`, id).Scan(
		&admin.Id,
		&admin.Full_Name,
		&admin.Email,
		&admin.Age,
		&admin.Status,
		&admin.Login,
		&admin.Password,
		&admin.Created_at); err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

func (c *adminRepo) Delete(id string) error {
	query := `delete from "admin" where id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
