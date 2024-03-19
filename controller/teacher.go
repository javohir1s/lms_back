package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Teacher(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateTeacher(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllTeacher(w, r)
		} else {
			c.GetByIDTeacher(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateTeacher(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteTeacher(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{}

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Teacher().Create(teacher)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{}
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	teacher.Id = r.URL.Query().Get("id")
	err := uuid.Validate(teacher.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Teacher().Update(teacher)
	if err != nil {
		fmt.Println("error while updating teacher,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllTeacher(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
		request = models.GetAllTeachersRequest{}
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	request.Search = search

	page, err := ParsePageQueryParam(r)
	if err != nil {
		fmt.Println("error while parsing page, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(r)
	if err != nil {
		fmt.Println("error while parsing limit, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit


	teachers, err := c.Store.Teacher().GetAll(request)
	if err != nil {
		fmt.Println("error while getting teachers,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, teachers)
}

func (c Controller) GetByIDTeacher(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	teacher, err := c.Store.Teacher().GetByID(id)
	if err != nil {
		fmt.Println("error while getting teacher by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, teacher)
}

func (c Controller) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Teacher().Delete(id)
	if err != nil {
		fmt.Println("error while deleting teacher, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
