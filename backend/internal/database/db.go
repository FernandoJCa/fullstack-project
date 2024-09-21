package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	_ "log"
	"time"

	"backend/internal/models"
)

var ErrNotFound = errors.New("record not found")

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) InitSchema() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func (db *DB) CreateTodo(todo *models.Todo) error {
	now := time.Now()
	res, err := db.Exec("INSERT INTO todos (title, completed, created_at) VALUES (?, ?, ?)", todo.Title, todo.Completed, now)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = id
	todo.CreatedAt = now
	return nil
}

func (db *DB) GetTodos() ([]models.Todo, error) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (db *DB) GetTodo(id int64) (models.Todo, error) {
	var t models.Todo
	err := db.QueryRow("SELECT id, title, completed, created_at FROM todos WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return t, ErrNotFound
	}
	if err != nil {
		return t, err
	}
	return t, nil
}

func (db *DB) UpdateTodo(todo *models.Todo) error {
	_, err := db.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteTodo(id int64) error {
	res, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}
