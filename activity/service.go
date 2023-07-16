package activity

type ActivityService interface{

	GetActivities()([]Activity, error)
}

type service struct {
	reporsitory ActivityRepository
}

func ImplActivityService(reporsitory ActivityRepository) *service {
	return &service{reporsitory}
}

func (s *service)GetActivities()([]Activity, error){
	activities, err := s.reporsitory.FindAll()

	if err != nil {
		return activities, err
	}

	return activities, nil
}