package user

import "gorm.io/gorm"

//type repository berhubungan nya dengan database langsung jadi di dalam
//interface repository ini terdapat fungsi untuk melakukan penyimpan/pengecekan/penhapusan dsb terhadap database
//inilah proses layering

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email= ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id= ?", ID).Find(&user).Error
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
