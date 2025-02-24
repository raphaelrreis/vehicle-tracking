package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Vehicle struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewVehicle(db *sql.DB) *Vehicle {
	return &Vehicle{db: db}
}

func (c *Vehicle) Create(name, description, categoryID string) (*Vehicle, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO vehicles (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Vehicle{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Vehicle) FindAll() ([]Vehicle, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM vehicles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vehicles := []Vehicle{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, Vehicle{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return vehicles, nil
}

func (c *Vehicle) FindByCategoryID(categoryID string) ([]Vehicle, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM vehicles WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vehicles := []Vehicle{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, Vehicle{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return vehicles, nil
}

func (c *Vehicle) Find(id string) (Vehicle, error) {
	var name, description, categoryID string
	err := c.db.QueryRow("SELECT name, description, category_id FROM vehicles WHERE id = $1", id).
		Scan(&name, &description, &categoryID)
	if err != nil {
		return Vehicle{}, err
	}
	return Vehicle{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}
