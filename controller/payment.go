package controller

import (
	"encoding/json"
	"fmt"
	"lms_back/models"

	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Payment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreatePayment(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllPayment(w, r)
		} else {
			c.GetByIDPayment(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdatePayment(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeletePayment(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method did not allowed")
	}
}

func (c Controller) CreatePayment(w http.ResponseWriter, r *http.Request) {
	payment := models.Payment{}

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	id, err := c.Store.Payment().Create(payment)
	if err != nil {
		fmt.Println("error while creating payment, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	payment := models.Payment{}
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	payment.Id = r.URL.Query().Get("id")
	err := uuid.Validate(payment.Id)
	if err != nil {
		fmt.Println("error while validating, err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Payment().Update(payment)
	if err != nil {
		fmt.Println("error while updating payment,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllPayment(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
		request = models.GetAllPaymentsRequest{}
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

	payment, err := c.Store.Payment().GetAll(request)
	if err != nil {
		fmt.Println("error while getting payment,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, payment)
}

func (c Controller) GetByIDPayment(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	payment, err := c.Store.Payment().GetByID(id)
	if err != nil {
		fmt.Println("error while getting payment by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, payment)
}

func (c Controller) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Payment().Delete(id)
	if err != nil {
		fmt.Println("error while deleting payment, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
