package gorm

import "gorm.io/gorm"

type Database struct {
	*gorm.DB
}

type IDatabase interface {
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Updates(values interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
}

type MockDatabase struct {
	FisrtFn   func(dest interface{}, conds ...interface{}) *gorm.DB
	CreateFn  func(value interface{}) *gorm.DB
	ModelFn   func(value interface{}) *gorm.DB
	UpdatesFn func(values interface{}) *gorm.DB
	DeleteFn  func(value interface{}, conds ...interface{}) *gorm.DB
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type IRepository interface {
	Get(id int) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}

type Repository struct {
	db IDatabase
}
