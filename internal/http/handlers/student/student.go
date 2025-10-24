package student

import (
    "encoding/json"
    "errors"
    "io"
    "log/slog"
    "net/http"
    "strconv"

    "github.com/CodeMaverick-143/Golang_learning/internal/storage"
    "github.com/CodeMaverick-143/Golang_learning/internal/types"
    "github.com/CodeMaverick-143/Golang_learning/internal/utils/response"
    "github.com/go-playground/validator/v10"
)

//Create 

func New(storage storage.Storage) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        slog.Info("Creating a new student")

        var student types.Student

        err := json.NewDecoder(r.Body).Decode(&student)
        if errors.Is(err, io.EOF) {
            _ = response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }
        if err != nil {
            _ = response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }

        if err := validator.New().Struct(student); err != nil {
            _ = response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
            return
        }

        lastID, err := storage.CreateStudent(
            student.Name,
            student.Email,
            student.Age,
        )
        if err != nil {
            _ = response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
            return
        }

        slog.Info("Student created successfully", slog.Int64("id", lastID))

        _ = response.WriteJSON(w, http.StatusCreated, map[string]string{
            "id": strconv.FormatInt(lastID, 10),
        })
    })
}

//Read

//Update

//Delete
