package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	_ "github.com/mattn/go-sqlite3"
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

	stmt , err := s.Db.Prepare(`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`)

	if err != nil {
		return 0, err
	}
	
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	
	lastId , err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId,nil
	
}
