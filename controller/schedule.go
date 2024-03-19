package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Schedule(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateSchedule(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllSchedule(w, r)
		} else {
			c.GetByIDSchedule(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateSchedule(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteSchedule(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	schedule := models.Schedule{}

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Schedule().Create(schedule)
	if err != nil {
		fmt.Println("error while creating schedule, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	schedule := models.Schedule{}
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	schedule.Id = r.URL.Query().Get("id")
	err := uuid.Validate(schedule.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Schedule().Update(schedule)
	if err != nil {
		fmt.Println("error while updating schedule,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllSchedule(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
		request = models.GetAllSchedulesRequest{}
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

	schedule, err := c.Store.Schedule().GetAll(request)
	if err != nil {
		fmt.Println("error while getting schedule,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, schedule)
}

func (c Controller) GetByIDSchedule(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	schedule, err := c.Store.Schedule().GetByID(id)
	if err != nil {
		fmt.Println("error while getting schedule by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, schedule)
}

func (c Controller) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Schedule().Delete(id)
	if err != nil {
		fmt.Println("error while deleting schedule, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
