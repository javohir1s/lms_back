package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"
	"lms_back/pkg"

	"github.com/google/uuid"
)

type TeacherRepo struct {
	db *sql.DB
}

func NewTeacher(db *sql.DB) TeacherRepo {
	return TeacherRepo{
		db: db,
	}
}

func (c *TeacherRepo) Create(teacher models.Teacher) (string, error) {

	id := uuid.New()
	query := `INSERT INTO teacher (
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
		teacher.Full_name,
	    teacher.Email,
		teacher.Age,
		teacher.Status,
		teacher.Login,
		teacher.Password,
	)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c *TeacherRepo) Update(teacher models.Teacher) (string, error) {
	query := `update teacher set 
	full_name=$1,
	email=$2,
	age=$3,
	login=$4,
	password=$5,
	status=$6,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $7 AND deleted_at = 0
	`
	_, err := c.db.Exec(query,
		teacher.Full_name,
		teacher.Email,
		teacher.Age,
		teacher.Login,
		teacher.Password,
		teacher.Status,
		teacher.Id,
    )
	if err != nil {
		return "", err
	}
	return teacher.Id, nil
}

func (c *TeacherRepo) GetAll(req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
    var (
		resp   = models.GetAllTeachersResponse{}
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
        login,
		password,
		status,
        created_at,
        updated_at,
        deleted_at FROM teacher WHERE deleted_at = 0` + filter + ``)
    if err != nil {
        return resp, err
    }
    for rows.Next() {
        var (
            teacher    = models.Teacher{}
            updateAt  sql.NullString
        )
        if err := rows.Scan(
            &resp.Count,
            &teacher.Id,
            &teacher.Full_name,
			&teacher.Email,
			&teacher.Age,
            &teacher.Login,
			&teacher.Password,
			&teacher.Status,
            &teacher.Created_at,
            &updateAt,
            &teacher.Deleted_at); err != nil {
            return resp, err
        }
        teacher.Updated_at = pkg.NullStringToString(updateAt)
        resp.Teachers = append(resp.Teachers, teacher)
    }
    return resp, nil
}

func (c *TeacherRepo) GetByID(id string) (models.Teacher, error) {
	teacher := models.Teacher{}

	if err := c.db.QueryRow(`select id, full_name, email, age, login, password, status, created_at from teacher where id = $1`, id).Scan(
		&teacher.Id,
		&teacher.Full_name,
		&teacher.Email,
		&teacher.Age,
		&teacher.Login,
		&teacher.Password,
		&teacher.Status,
		&teacher.Created_at,
		); err != nil {
		return models.Teacher{}, err
	}
	return teacher, nil
}

func (c *TeacherRepo) Delete(id string) error {
	query := `delete from teacher where id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}