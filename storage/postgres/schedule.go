package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"
	"lms_back/pkg"

	"github.com/google/uuid"
)

type ScheduleRepo struct {
	db *sql.DB
}

func NewSchedule(db *sql.DB) ScheduleRepo {
	return ScheduleRepo{
		db: db,
	}
}

func (c *ScheduleRepo) Create(schedule models.Schedule) (string, error) {

	id := uuid.New()
	query := `INSERT INTO schedule (
		id,
		group_id,
		group_type,
		start_time,
		end_time,
		date,
		branch_id,
		teacher_id,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(query,
		id.String(),
		schedule.Group_id,
	    schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id,
		schedule.Created_at,
	)

	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c *ScheduleRepo) Update(schedule models.Schedule) (string, error) {
	query := `update schedule set 
	group_id=$1,
	group_type=$2,
	start_time=$3,
	end_time=$4,
	date=$5,
	branch_id=$6,
	teacher_id=$7,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $8
	`
	_, err := c.db.Exec(query,
		schedule.Group_id,
	    schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id,
		schedule.Created_at,
		schedule.Id,
    )
	if err != nil {
		return "", err
	}
	return schedule.Id, nil
}

func (c *ScheduleRepo) GetAll(req models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error) {
    var (
		resp   = models.GetAllSchedulesResponse{}
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
		group_id,
		group_type,
		start_time,
		end_time,
		date,
		branch_id,
		teacher_id,
		created_at,
        updated_at FROM schedule` + filter + ``)
	
    if err != nil {
        return resp, err
    }

    for rows.Next() {
        var (
            schedule   = models.Schedule{}
            updateAt  sql.NullString
			// group_id  sql.NullString
        )
        if err := rows.Scan(
            &resp.Count,
            &schedule.Id,
            &schedule.Group_id,
			&schedule.Group_type,
			&schedule.Start_time,
			&schedule.End_time,
			&schedule.Date,
            &schedule.Branch_id,
			&schedule.Teacher_id,
			&schedule.Created_at,
            &updateAt);
			 err != nil {
            return resp, err
        }
        schedule.Updated_at = pkg.NullStringToString(updateAt)
        resp.Schedules = append(resp.Schedules, schedule)
    }
    return resp, nil
}

func (c *ScheduleRepo) GetByID(id string) (models.Schedule, error) {
	schedule := models.Schedule{}
	// var (
	// 	group_id  sql.NullString
	// )
	if err := c.db.QueryRow(`	
			select 
			id, 
			group_id,
			group_type,
			start_time,
			end_time,
			date,
			branch_id,
			teacher_id,
			created_at,
			updated_at from schedule where id = $1`, id).Scan(
			schedule.Group_id,
			schedule.Group_type,
			schedule.Start_time,
			schedule.End_time,
			schedule.Date,
			schedule.Branch_id,
			schedule.Teacher_id,
			schedule.Created_at,
		); err != nil {
		return models.Schedule{}, err
	}
	return schedule, nil
}


func (c *ScheduleRepo) Delete(id string) error {
	query := `delete from schedule where id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}