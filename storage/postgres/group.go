package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"
	"lms_back/pkg"
	"strconv"

	"github.com/google/uuid"
)

type GroupRepo struct {
	db *sql.DB
}

func NewGroup(db *sql.DB) GroupRepo {
	return GroupRepo{
		db: db,
	}
}

func (g *GroupRepo) Create(group models.Group) (string, error) {
	id := uuid.New()
	var group_unique_id string

	maxQuery := `
	SELECT MAX(group_id) 
	FROM "group"
	`
	err := g.db.QueryRow(maxQuery).Scan(&group_unique_id)
	if err != nil {
		if err.Error() != "can't scan into dest[0]: cannot scan null into *string" && err.Error() != "no rows in result set" {
			return "resp", err
		} else {
			group_unique_id = "GR-0000000"
		}
	}

	digit := 0
	if len(group_unique_id) > 2 {
		digit, err = strconv.Atoi(group_unique_id[3:])
		if err != nil {
			return "resp", err
		}
	}

	query := `INSERT INTO "group" (
		id,
		group_id,
		branch_id,
		teacher_id,
		type,
		created_at) 
		VALUES($1,$2,$3,$4,$5,CURRENT_TIMESTAMP)
		`

	_, err = g.db.Exec(query,
		id.String(),
		"Gr-" + pkg.GetSerialId(digit),
		group.Branch_id,
		group.Teacher_id,
		group.Type)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (g *GroupRepo) Update(group models.Group) (string, error) {
	query := `UPDATE "group" SET
		type=$1,
		updated_at=CURRENT_TIMESTAMP
		WHERE id=$2`

	_, err := g.db.Exec(query, group.Type, group.Id)
	if err != nil {
		return "", nil
	}
	return group.Id, nil
}

func (g *GroupRepo) GetAll(req models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {
	var (
		resp   = models.GetAllGroupsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := g.db.Query(`SELECT count (id) OVER(),
        id,
        group_id,
        branch_id,
        teacher_id,
        type,
        created_at,
        updated_at FROM "group"` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			group      = models.Group{}
			updateAt   sql.NullString
			teacher_id sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&group.Id,
			&group.Group_id,
			&group.Branch_id,
			&teacher_id,
			&group.Type,
			&group.Created_at,
			&updateAt); err != nil {
			return resp, err
		}
		group.Updated_at = pkg.NullStringToString(updateAt)
		group.Teacher_id = pkg.NullStringToString(teacher_id)
		resp.Groups = append(resp.Groups, group)
	}
	return resp, nil
}

func (g *GroupRepo) GetByID(id string) (models.Group, error) {
	group := models.Group{}

	if err := g.db.QueryRow(`SELECT id, group_id, branch_id, teacher_id, type, created_at, updated_at FROM "group" WHERE id = $1`, id).Scan(
		&group.Id,
		&group.Group_id,
		&group.Branch_id,
		&group.Teacher_id,
		&group.Type,
		&group.Created_at,
		&group.Updated_at); err != nil {
		return models.Group{}, err
	}
	return group, nil
}

func (g *GroupRepo) Delete(id string) error {
	query := `delete from "group" where id = $1`
	_, err := g.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
