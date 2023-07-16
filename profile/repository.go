package profile

import (
	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindAll() ([]Profile, error)
	FindByID(profileID int) ([]Profile, error)
}

type reporsitory struct {
	db *gorm.DB
}

func ImplProfileRepository(db *gorm.DB) *reporsitory{
	return &reporsitory{db} // initiate db di panggil di service
}

func (r *reporsitory) FindAll() ([]Profile, error) {
	var profiles []Profile

	err := r.db.Find(&profiles).Error

	if err != nil {
		return profiles, err
	}

	return profiles, nil
}

func (r *reporsitory) FindByID(profileID int) ([]Profile, error){
	var profiles []Profile
	err := r.db.Where("id = ?", profileID).Find(&profiles).Error

	if err != nil {
		return profiles, nil
	} 
	return profiles, nil
}