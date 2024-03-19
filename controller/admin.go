package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Admin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateAdmin(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllAdmins(w, r)
		} else {
			c.GetByIDAdmin(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateAdmin(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteAdmin(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	admin := models.Admin{}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Admin().Create(admin)
	if err != nil {
		fmt.Println("error while creating admin, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	admin := models.Admin{}
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	admin.Id = r.URL.Query().Get("id")
	err := uuid.Validate(admin.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Admin().Update(admin)
	if err != nil {
		fmt.Println("error while updating admin, err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	var (
		values  = r.URL.Query()
		search  string
		request = models.GetAllAdminsRequest{}
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

	admins, err := c.Store.Admin().GetAll(request)
	if err != nil {
		fmt.Println("error while getting admins, err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, admins)
}

func (c Controller) GetByIDAdmin(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	admin, err := c.Store.Admin().GetByID(id)
	if err != nil {
		fmt.Println("error while getting admin by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, admin)
}

func (c Controller) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Admin().Delete(id)
	if err != nil {
		fmt.Println("error while deleting admin, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
