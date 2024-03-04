package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return nil, err
	}

	return &Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var id, name, description string

		err := rows.Scan(&id, &name, &description)

		if err != nil {
			return nil, err
		}

		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {
	var id, name, description string

	err := c.db.QueryRow("SELECT ct.id, ct.name, ct.description FROM categories ct JOIN courses cs ON cs.category_id = ct.id WHERE cs.id = $1", courseId).
		Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, err
}

func (c *Category) FindById(id string) (*Category, error) {
	var name, description string

	err := c.db.QueryRow("SELECT name, description FROM categories WHERE id = $1", id).Scan(&name, &description)

	if err != nil {
		return nil, err
	}

	return &Category{ID: id, Name: name, Description: description}, nil
}
