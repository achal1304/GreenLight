package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
		GetAll(Title string, Genres []string, Filters Filters) ([]*Movie, Metadata, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{Movies: MovieModel{DB: db}}
}

func NewMockModels() Models {
	return Models{Movies: MockMovieModel{}}
}
