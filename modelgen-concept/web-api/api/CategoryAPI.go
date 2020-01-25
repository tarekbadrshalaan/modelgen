package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dto"
)

func configCategoriesRouter(router *httprouter.Router) {
	router.GET("/categories", getAllCategories)
	router.POST("/categories", postCategories)
	router.PUT("/categories", putCategories)
	router.GET("/categories/:id", getCategories)
	router.DELETE("/categories/:id", deleteCategories)
}

func getAllCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categories, err := bll.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, categories)
}

func getCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCategoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	category, err := bll.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, category)
}


func postCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCategoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

