package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	"github.com/CodeMaverick-143/Golang_learning/internal/storage"
	_ "github.com/mattn/go-sqlite3"
	"github.com/CodeMaverick-143/Golang_learning/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	dbPath, err := filepath.Abs(cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare(`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`)

	if err != nil {
		return 0, err
	}
	
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
	
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT id , name , email , age FROM students WHERE id = ? LIMIT 1`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	row := stmt.QueryRow(id)

	if err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Age); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.Student{}, storage.ErrNotFound
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}
