package gorm

import "gorm.io/gorm"

func (m MockDatabase) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return m.FisrtFn(dest, conds...)
}

func (m MockDatabase) Create(value interface{}) *gorm.DB {
	return m.CreateFn(value)
}

func (m MockDatabase) Model(value interface{}) *gorm.DB {
	return m.ModelFn(value)
}

func (m MockDatabase) Updates(values interface{}) *gorm.DB {
	return m.UpdatesFn(values)
}

func (m MockDatabase) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return m.DeleteFn(value, conds...)
}
