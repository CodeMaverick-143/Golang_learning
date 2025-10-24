package storage

import (
    "errors"

    "github.com/CodeMaverick-143/Golang_learning/internal/types"
)

var ErrNotFound = errors.New("storage: not found")

type Storage interface {
    CreateStudent(name string, email string, age int) (int64, error)
    GetStudentById(id int64) (types.Student, error)
}