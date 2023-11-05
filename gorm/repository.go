package gorm

func NewRepository(db IDatabase) *Repository {
	return &Repository{db}
}
func (d *Repository) Get(id uint) (*Product, error) {
	product := &Product{}
	res := d.db.First(&product, id)

	return product, res.Error
}

func (d *Repository) Create(product *Product) error {
	return d.db.Create(product).Error
}

func (d *Repository) Update(product *Product) error {
	oldProduct, err := d.Get(product.ID)
	if err != nil {
		return err
	}

	return d.db.Model(oldProduct).Updates(product).Error
}

func (d *Repository) Delete(id uint) error {
	product, err := d.Get(id)
	if err != nil {
		return err
	}

	return d.db.Delete(product).Error
}
