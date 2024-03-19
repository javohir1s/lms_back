package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Branch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBranch(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllBranches(w, r)
		} else {
			c.GetByIDBranch(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateBranch(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteBranch(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreateBranch(w http.ResponseWriter, r *http.Request) {
	branch := models.Branches{}

	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Branches().Create(branch)
	if err != nil {
		fmt.Println("error while creating branch, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	branch := models.Branches{}
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	branch.Id = r.URL.Query().Get("id")
	err := uuid.Validate(branch.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Branches().Update(branch)
	if err != nil {
		fmt.Println("error while updating branch,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllBranches(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
		request = models.GetAllBranchesRequest{}
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

	branches, err := c.Store.Branches().GetAll(request)
	if err != nil {
		fmt.Println("error while getting branches,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, branches)
}

func (c Controller) GetByIDBranch(w http.ResponseWriter,r *http.Request)  {
	values:= r.URL.Query()
	id :=values["id"][0]

	customer,err := c.Store.Branches().GetByID(id)
	if err != nil {
		fmt.Println("error while getting branch by id")
		handleResponse(w,http.StatusInternalServerError,err)
		return
	}
	handleResponse(w,http.StatusOK,customer)
}

func (c Controller) DeleteBranch(w http.ResponseWriter,r *http.Request)  {
	id := r.URL.Query().Get("id")
	fmt.Println("id",id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:",err.Error())
	handleResponse(w,http.StatusBadRequest,err.Error())
	return
}
err = c.Store.Branches().Delete(id)
if err != nil {
	fmt.Println("error while deleting branch, err:",err)
	handleResponse(w,http.StatusInternalServerError,err)
return
}
handleResponse(w,http.StatusOK,id)
}