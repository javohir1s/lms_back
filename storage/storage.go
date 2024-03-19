package storage

import (
	"lms_back/models"
)

type IStorage interface {
	CloseDB()
	Admin() IAdminStorage
	Branches() IBranchStorage
	Group() IGroupStorage
	Payment() IPaymentStorage
	Student() IStudentStorage
	Teacher() ITeacherStorage
	Schedule() IScheduleStorage
}

type IAdminStorage interface {
	Create(models.Admin) (string, error)
	GetAll(request models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(id string) (models.Admin, error)
	Update(models.Admin) (string, error)
	Delete(string) error
}

type IBranchStorage interface {
	Create(models.Branches) (string, error)
	GetAll(request models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(id string) (models.Branches, error)
	Update(models.Branches) (string, error)
	Delete(string) error
}

type IGroupStorage interface {
	Create(models.Group) (string, error)
	GetAll(request models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(id string) (models.Group, error)
	Update(models.Group) (string, error)
	Delete(string) error
}

type IPaymentStorage interface {
	Create(models.Payment) (string, error)
	GetAll(request models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(id string) (models.Payment, error)
	Update(models.Payment) (string, error)
	Delete(string) error
}

type IStudentStorage interface {
	Create(models.Student) (string, error)
	GetAll(request models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(id string) (models.Student, error)
	Update(models.Student) (string, error)
	Delete(string) error
}

type ITeacherStorage interface {
	Create(models.Teacher) (string, error)
	GetAll(request models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(id string) (models.Teacher, error)
	Update(models.Teacher) (string, error)
	Delete(string) error
}

type IScheduleStorage interface {
	Create(models.Schedule) (string, error)
	GetAll(request models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error)
	GetByID(id string) (models.Schedule, error)
	Update(models.Schedule) (string, error)
	Delete(string) error
}
