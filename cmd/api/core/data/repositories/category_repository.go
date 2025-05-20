package repositories

import (
	"database/sql"
	"sync"
)

type CategoryRepository struct {
	mutex sync.Mutex
	db    *sql.DB
}

func (r *CategoryRepository) FindById() {}
func (r *CategoryRepository) Create()   {}
func (r *CategoryRepository) List()     {}
func (r *CategoryRepository) Products() {}
func (r *CategoryRepository) Delete()   {}
