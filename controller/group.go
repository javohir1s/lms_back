package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Group(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateGroup(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllGroups(w, r)
		} else {
			c.GetByIDGroup(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateGroup(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteGroup(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreateGroup(w http.ResponseWriter, r *http.Request) {
	group := models.Group{}

	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Group().Create(group)
	if err != nil {
		fmt.Println("error while creating group, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	group := models.Group{}
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	group.Id = r.URL.Query().Get("id")
	err := uuid.Validate(group.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Group().Update(group)
	if err != nil {
		fmt.Println("error while updating group,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
		request = models.GetAllGroupsRequest{}
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

	groups, err := c.Store.Group().GetAll(request)
	if err != nil {
		fmt.Println("error while getting groups,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, groups)
}

func (c Controller) GetByIDGroup(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	group, err := c.Store.Group().GetByID(id)
	if err != nil {
		fmt.Println("error while getting group by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, group)
}

func (c Controller) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Group().Delete(id)
	if err != nil {
		fmt.Println("error while deleting group, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
