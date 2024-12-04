package data

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Year      int       `json:"year,omitempty"`
	Runtime   int       `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	CreatedAt time.Time `json:"-"`
	Version   int       `json:"version"`
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

func (m MovieModel) Get(id int) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int) error {
	return nil
}
