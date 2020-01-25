package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configCustomersRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/customers", handle: getAllCustomers})
	*routes = append(*routes, route{method: "POST", path:"/customers", handle: postCustomers})
	*routes = append(*routes, route{method: "PUT", path:"/customers", handle: putCustomers})
	*routes = append(*routes, route{method: "GET", path:"/customers/:id", handle: getCustomers})
	*routes = append(*routes, route{method: "DELETE", path:"/customers/:id", handle: deleteCustomers})
}

func getAllCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customers, err := bll.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, customers)
}

func getCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCustomerID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	customer, err := bll.GetCustomer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, customer)
}

func postCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customer := &dto.CustomerDTO{}
	if err := readJSON(r, customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customer := &dto.CustomerDTO{}
	if err := readJSON(r, customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func deleteCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCustomerID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteCustomer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
