package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/models"

	"github.com/google/uuid"
)

type paymentRepo struct {
	db *sql.DB
}

func NewPayment(db *sql.DB) paymentRepo {
	return paymentRepo{
		db: db,
	}
}

func (p *paymentRepo) Create(payment models.Payment) (string, error) {
	id := uuid.New()

	query := `INSERT INTO payment(id, price, student_id, branch_id, admin_id, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),CREATE OR REPLACE FUNCTION update_student_paid_sum()
			  RETURNS TRIGGER AS $$
			  BEGIN
				  IF TG_OP = 'INSERT' THEN
					  UPDATE "student"
					  SET "paid_sum" = "paid_sum" + NEW.price
					  WHERE "id" = NEW.student_id;
				  END IF;
				  RETURN NEW;
			  END;
			  $$ LANGUAGE plpgsql;
			  
			  CREATE TRIGGER update_student_paid_sum_trigger
			  AFTER INSERT ON "payment"
			  FOR EACH ROW
			  EXECUTE FUNCTION update_student_paid_sum();
			  `

	_, err := p.db.Exec(query, id.String(), payment.Price, payment.Student_id, payment.Branch_id, payment.Admin_id)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (p *paymentRepo) GetAll(req models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error) {
	var (
		resp   = models.GetAllPaymentsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := p.db.Query(`SELECT id, price, student_id, branch_id, admin_id, created_at, updated_at FROM payment`)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var payment models.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.Price,
			&payment.Student_id,
			&payment.Branch_id,
			&payment.Admin_id,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		); err != nil {
			return resp, err
		}
		resp.Payments = append(resp.Payments, payment)
	}
	return resp, nil
}

func (p *paymentRepo) GetByID(id string) (models.Payment, error) {
	var payment models.Payment

	row := p.db.QueryRow(`SELECT id, price, student_id, branch_id, admin_id, created_at, updated_at FROM payment WHERE id = $1`, id)
	if err := row.Scan(
		&payment.Id,
		&payment.Price,
		&payment.Student_id,
		&payment.Branch_id,
		&payment.Admin_id,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	); err != nil {
		return payment, err
	}

	return payment, nil
}

func (p *paymentRepo) Update(payment models.Payment) (string, error) {
	query := `UPDATE payment SET price=$1, student_id=$2, branch_id=$3, admin_id=$4, updated_at=CURRENT_TIMESTAMP WHERE id=$5,
	CREATE OR REPLACE FUNCTION update_student_paid_sum()
	RETURNS TRIGGER AS $$
	BEGIN
		IF TG_OP = 'INSERT' THEN
			UPDATE "student"
			SET "paid_sum" = "paid_sum" + NEW.price
			WHERE "id" = NEW.student_id;
		END IF;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	
	CREATE TRIGGER update_student_paid_sum_trigger
	AFTER Update ON "payment"
	FOR EACH ROW
	EXECUTE FUNCTION update_student_paid_sum();
	`
	_, err := p.db.Exec(query, payment.Price, payment.Student_id, payment.Branch_id, payment.Admin_id, payment.Id)
	if err != nil {
		return "", err
	}
	return payment.Id, nil
}

func (p *paymentRepo) Delete(id string) error {
	_, err := p.db.Exec(`DELETE FROM payment WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
