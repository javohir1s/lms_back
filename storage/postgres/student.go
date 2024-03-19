package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"
	"lms_back/pkg"

	"github.com/google/uuid"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) StudentRepo {
	return StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()
	query := `INSERT INTO student (
		id,
		full_name,
		email,
		age,
		paid_sum,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,CURRENT_TIMESTAMP) 
	`
	// CREATE TABLE IF NOT EXISTS "student" (
	// 	"id" uuid PRIMARY KEY,
	// 	"full_name" varchar(255) NOT NULL,
	// 	"email" varchar(255) NOT NULL,
	// 	"age" int NOT NULL,
	// 	"paid_sum" decimal(10, 2) NOT NULL DEFAULT 0,
	// 	"status" varchar(60) NOT NULL CHECK("status" IN ('active', 'inactive')) DEFAULT 'active',
	// 	"login" varchar(255) NOT NULL,
	// 	"password" varchar(255) NOT NULL,
	// 	"group_id" uuid REFERENCES "group"("id"),
	// 	"created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// 	"updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	//   );
	_, err := c.db.Exec(query,
		id.String(),
		student.Full_Name,
	    student.Email,
		student.Age,
		student.PaidSum,
		student.Status,
		student.Login,
		student.Password,
	)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c *StudentRepo) Update(student models.Student) (string, error) {
	query := `update student set 
	full_name=$1,
	email=$2,
	age=$3,
	paid_sum=$4,
	login=$5,
	password=$6,
	group_id=$7,
	status=$8,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $9
	`
	_, err := c.db.Exec(query,
		student.Full_Name,
		student.Email,
		student.Age,
		student.PaidSum,
		student.Login,
		student.Password,
		student.GroupID,
		student.Status,
		student.ID,
    )
	if err != nil {
		return "", err
	}
	return student.ID, nil
}

func (c *StudentRepo) GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
    var (
		resp   = models.GetAllStudentsResponse{}
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
		paid_sum,
		status,
        login,
		password,
		group_id,
        created_at,
        updated_at FROM student` + filter + ``)
    if err != nil {
        return resp, err
    }

    for rows.Next() {
        var (
            student   = models.Student{}
            updateAt  sql.NullString
			// group_id  sql.NullString
        )
        if err := rows.Scan(
            &resp.Count,
            &student.ID,
            &student.Full_Name,
			&student.Email,
			&student.Age,
			&student.PaidSum,
			&student.Status,
            &student.Login,
			&student.Password,
			&student.GroupID,
            &student.Created_At,
            &updateAt);
			 err != nil {
            return resp, err
        }
        student.Updated_At = pkg.NullStringToString(updateAt)
        resp.Students = append(resp.Students, student)
    }
    return resp, nil
}

func (c *StudentRepo) GetByID(id string) (models.Student, error) {
	student := models.Student{}
	// var (
	// 	group_id  sql.NullString
	// )
	if err := c.db.QueryRow(`select id, full_name, email, age, paid_sum, status, login, password, group_id, created_at from student where id = $1`, id).Scan(
		&student.ID,
		&student.Full_Name,
		&student.Email,
		&student.Age,
		&student.PaidSum,
		&student.Status,
		&student.Login,
		&student.Password, 
		&student.GroupID,
		&student.Created_At,
		); err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (c *StudentRepo) Delete(id string) error {
	query := `delete from student where id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}