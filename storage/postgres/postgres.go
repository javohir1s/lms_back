package postgres

import (
	"database/sql"
	"fmt"
	"lms_back/config"
	"lms_back/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return Store{
		DB: db,
	}, nil

}

func (s Store) CloseDB() {
	s.DB.Close()
}

func (s Store) Admin() storage.IAdminStorage {
	NewAdmin := NewAdmin(s.DB)

	return &NewAdmin
}

func (s Store) Branches() storage.IBranchStorage {
	NewBranch := NewBranch(s.DB)

	return &NewBranch
}

func (s Store) Group() storage.IGroupStorage {
	NewGroup := NewGroup(s.DB)

	return &NewGroup
}

func (s Store) Payment() storage.IPaymentStorage {
	NewPayment := NewPayment(s.DB)

	return &NewPayment
}

func (s Store) Student() storage.IStudentStorage {
	NewStudent := NewStudent(s.DB)

	return &NewStudent
}

func (s Store) Teacher() storage.ITeacherStorage {
	NewTeacher := NewTeacher(s.DB)

	return &NewTeacher
}

func (s Store) Schedule() storage.IScheduleStorage {
	NewSchedule := NewSchedule(s.DB)

	return &NewSchedule
}