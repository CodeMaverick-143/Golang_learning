package student

import (
	"net/http"
	"log/slog"
	"encoding/json"
	"errors"
	"io"
	"github.com/CodeMaverick-143/Golang_learning/internal/types"
	"github.com/CodeMaverick-143/Golang_learning/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

//Create 

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a new student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}


		if err != nil{
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation

		if err := validator.New().Struct(student); err != nil{
			response.WriteJSON(w, http.StatusBadRequest,response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		response.WriteJSON(w, http.StatusCreated, map[string]string{"success":"OK"})
	})
}

//Read

//Update

//Delete
