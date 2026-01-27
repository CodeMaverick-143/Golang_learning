package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	"github.com/CodeMaverick-143/Golang_learning/internal/storage"
	"github.com/CodeMaverick-143/Golang_learning/internal/types"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Postgres, error) {
	db, err := sql.Open("pgx", cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}

func (s *Postgres) CreateStudent(name string, email string, age int) (int64, error) {
	var lastId int64
	err := s.Db.QueryRow(`INSERT INTO students (name, email, age) VALUES ($1, $2, $3) RETURNING id`, name, email, age).Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (s *Postgres) GetStudentById(id int64) (types.Student, error) {
	var student types.Student
	err := s.Db.QueryRow(`SELECT id, name, email, age FROM students WHERE id = $1`, id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.Student{}, storage.ErrNotFound
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil
}

func (s *Postgres) UpdateStudent(id int64, name string, email string, age int) error {
	_, err := s.Db.Exec(`UPDATE students SET name = $1, email = $2, age = $3 WHERE id = $4`, name, email, age, id)
	return err
}

func (s *Postgres) DeleteStudent(id int64) error {
	_, err := s.Db.Exec(`DELETE FROM students WHERE id = $1`, id)
	return err
}
