package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *Repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) UpdateByID(id uint, data *User) (*User, error) {
	user, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	user.Email = data.Email
	user.Password = data.Password
	err = r.db.Save(user).Error
	return user, err
}

func (r *Repository) DeleteByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
