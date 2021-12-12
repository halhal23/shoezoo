package repository

import "github.com/halhal23/shoezoo/product/domain/model"

type IProductRepo interface {
	Insert(*model.Product) (int, error)
	// GetAll() ([]*model.Product, error)
	// GetById(int) (*model.Product, error)
	// GetByName(string) (*model.Product, error)
	// Update(*model.Product) (*model.Product, error)
	// Delete(int) (int, error)
}
