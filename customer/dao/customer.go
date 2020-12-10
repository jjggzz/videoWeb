package dao

func (d *Dao) QueryCustomerByAccessKey(accessKey string) (*Customer, error) {
	customer := &Customer{}
	d.db.Where("access_key = ?", accessKey).First(customer)
	return customer, nil
}

func (d *Dao) InsertCustomer(customer *Customer) error {
	return d.db.Create(customer).Error
}
