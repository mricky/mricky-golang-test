package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(userId int) (User, error)
}

type reporsitory struct {
	db *gorm.DB
}

func ImplUserRepository(db *gorm.DB) *reporsitory{
	return &reporsitory{db} 
}

func (r *reporsitory) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *reporsitory) FindByEmail(email string) (User, error){
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, nil
	} 
	return user, nil
}

func (r *reporsitory) FindByID(userId int) (User, error){
	var user User
	err := r.db.Where("id = ?", userId).Find(&user).Error

	if err != nil {
		return user, nil
	} 
	return user, nil
}
