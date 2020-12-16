package dao

func (d *Dao) InsertCustomer(customer *Customer) error {
	return d.db.Create(customer).Error
}

func (d *Dao) ExitsCustomerByPhone(phone string) (bool, error) {
	var count int
	err := d.db.Model(&Customer{}).Where("phone = ?", phone).Count(&count).Error
	return count > 0, err
}

func (d *Dao) SelectCustomerByPhone(phone string) (*Customer, error) {
	customer := &Customer{}
	err := d.db.Where("phone = ?", phone).First(customer).Error
	return customer, err
}
