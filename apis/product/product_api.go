package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/keveaux/go_CRUD_application/entities"

	"github.com/gorilla/mux"

	"github.com/keveaux/go_CRUD_application/config"
	"github.com/keveaux/go_CRUD_application/models"
)

func FindAll(response http.ResponseWriter, request *http.Request) {

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.FindAll()

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	keyword := vars["keyword"]

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.Search(keyword)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func SearchPrices(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	smin := vars["min"]
	smax := vars["max"]

	min, _ := strconv.ParseFloat(smin, 64)
	max, _ := strconv.ParseFloat(smax, 64)

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.SearchPrices(min, max)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {

	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		err2 := productModel.Create(&product)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {

	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		_, err2 := productModel.Update(&product)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	sid := vars["id"]

	id, _ := strconv.ParseInt(sid, 10, 64)

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		productModel := models.ProductModel{
			Db: db,
		}

		_, err2 := productModel.Delete(id)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
