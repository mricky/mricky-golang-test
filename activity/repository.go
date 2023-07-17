package activity

import "gorm.io/gorm"

type ActivityRepository interface {
	Save(activity Activity) (Activity, error)
	Update(activity Activity) (Activity, error)
	FindAll()([]Activity, error)

}

type reporsitory struct {
	db *gorm.DB
}

func ImplActivityRepository(db *gorm.DB) *reporsitory {
	return &reporsitory{db}
}

func (r *reporsitory) Save(activity Activity) (Activity, error) {
	err := r.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}
	return activity, nil
}


func (r *reporsitory) Update(activity Activity) (Activity, error) {
	err := r.db.Save(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *reporsitory) FindAll() ([]Activity, error) {
	var activities []Activity

	err := r.db.Find(&activities).Error

	if err != nil {
		return activities, err
	}

	return activities, nil
}