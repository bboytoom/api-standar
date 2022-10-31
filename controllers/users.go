package controllers

import (
	"api_standard/helpers"
	"api_standard/models"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func Index(w http.ResponseWriter, r *http.Request) {
	retrieve_users, err := json.Marshal(models.RetrieveUsers())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(retrieve_users)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var user models.User
	validate = validator.New()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errorPrint("Error", err.Error(), nil, http.StatusInternalServerError, w)
		return
	}

	err_validate := validate.Struct(user)

	if err_validate != nil {
		errors_array := []helpers.ResponseError{}

		if _, ok := err_validate.(*validator.InvalidValidationError); ok {
			return
		}

		for _, err := range err_validate.(validator.ValidationErrors) {
			errors_array = append(errors_array, helpers.ResponseError{
				Field: err.StructField(),
				Type:  err.ActualTag(),
			})
		}

		errorPrint("Error", "", errors_array, http.StatusUnprocessableEntity, w)

		return
	}

	models.AddOnlyUser(user)

	result := models.ReturnUser{
		Result: "Added",
		Id:     1,
	}

	return_result, err := json.Marshal(result)

	if err != nil {
		errorPrint("Error", err.Error(), nil, http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(return_result)
}

func errorPrint(msg string, msg_error string, array_errors []helpers.ResponseError, status int, w http.ResponseWriter) {
	result := helpers.ErrorsGeneral{
		Result:  msg,
		Message: msg_error,
		Errors:  array_errors,
	}

	error_result, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(error_result)
}
