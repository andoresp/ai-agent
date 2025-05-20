package repositories

import (
	"database/sql"
	"sync"
)

type ProductRepository struct {
	mutex sync.Mutex
	db    *sql.DB
}

func (r *ProductRepository) FindById() {}
func (r *ProductRepository) Create()   {}
func (r *ProductRepository) List()     {}
func (r *ProductRepository) Delete()   {}
