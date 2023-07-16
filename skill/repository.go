package skill

import (
	"gorm.io/gorm"
)

type SkillRepository interface {
	FindAll() ([]Skill, error)
	FindByID(skilID int) ([]Skill, error)
}

type reporsitory struct {
	db *gorm.DB
}

func ImplSkillRepository(db *gorm.DB) *reporsitory{
	return &reporsitory{db} 
}

func (r *reporsitory) FindAll() ([]Skill, error) {
	var skills []Skill

	err := r.db.Find(&skills).Error

	if err != nil {
		return skills, err
	}

	return skills, nil
}

func (r *reporsitory) FindByID(skillID int) ([]Skill, error){
	var skills []Skill
	err := r.db.Where("id = ?", skillID).Find(&skills).Error

	if err != nil {
		return skills, nil
	} 
	return skills, nil
}