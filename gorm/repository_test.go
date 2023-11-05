package gorm

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// create table product (id integer primary key autoincrement, code text, price integer);
func Setup() {
	db := NewConnection("test.db")
	db.AutoMigrate(&Product{})
}

var ID uint = 5

//test integration with sqlite

func TestSqliteGorm(t *testing.T) {

	t.Run("Create", func(t *testing.T) {
		Setup()
		db := NewConnection("test.db")

		repo := NewRepository(db)

		product := &Product{
			Code:  "L1212",
			Price: 1000,
		}

		err := repo.Create(product)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		ID = product.ID
	})

	t.Run("Get", func(t *testing.T) {
		Setup()
		db := NewConnection("test.db")
		repo := NewRepository(db)
		product, err := repo.Get(ID)
		if err != nil {
			t.Fatalf("Failed to get product: %v", err)
		}

		if product.Code != "L1212" {
			t.Fatalf("Product code is not L1212: %v", product.Code)
		}
	})

	t.Run("Update", func(t *testing.T) {
		db := NewConnection("test.db")
		repo := NewRepository(db)
		product := &Product{
			Model: gorm.Model{
				ID: ID,
			},
			Code:  "L1212",
			Price: 2000,
		}

		err := repo.Update(product)
		if err != nil {
			t.Fatalf("Failed to update product: %v", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		db := NewConnection("test.db")
		repo := NewRepository(db)
		err := repo.Delete(ID) //soft delete
		if err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}
	})
}

func TestSqliteGormCaseError(t *testing.T) {
	t.Run("Create", func(t *testing.T) {

		db := MockDatabase{
			CreateFn: func(value interface{}) *gorm.DB {
				return &gorm.DB{
					Error: fmt.Errorf("error"),
				}
			},
		}

		repo := NewRepository(db)

		product := &Product{
			Code:  "L1212",
			Price: 1000,
		}

		err := repo.Create(product)
		if err == nil {
			t.Fatalf("Failed to create product: %v", err)
		}
	})

	t.Run("First", func(t *testing.T) {

		db := MockDatabase{
			FisrtFn: func(dest interface{}, conds ...interface{}) *gorm.DB {
				return &gorm.DB{
					Error: fmt.Errorf("error"),
				}
			},
		}

		repo := NewRepository(db)

		_, err := repo.Get(4)
		if err == nil {
			t.Fatalf("Failed to get product: %v", err)
		}
	})

	t.Run("Update", func(t *testing.T) {

		db := MockDatabase{
			FisrtFn: func(dest interface{}, conds ...interface{}) *gorm.DB {
				return &gorm.DB{
					Error: fmt.Errorf("error"),
				}
			},
		}

		repo := NewRepository(db)

		product := &Product{
			Model: gorm.Model{
				ID: 4,
			},
			Code:  "L1212",
			Price: 2000,
		}

		err := repo.Update(product)
		if err == nil {
			t.Fatalf("Failed to update product: %v", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {

		db := MockDatabase{
			FisrtFn: func(dest interface{}, conds ...interface{}) *gorm.DB {
				return &gorm.DB{
					Error: fmt.Errorf("error"),
				}
			},
		}

		repo := NewRepository(db)

		err := repo.Delete(4)
		if err == nil {
			t.Fatalf("Failed to delete product: %v", err)
		}
	})
}
