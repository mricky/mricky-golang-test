package activity

import "fmt"

type ActivityService interface{
	Create(input CreateActivityInput) (Activity,error)
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

func (s *service) Create(input CreateActivityInput)(Activity, error) {
	activity := Activity{}
	activity.SkillId = input.SkillId
	activity.UserId = input.UserId
	activity.StartDate = input.StartDate
	activity.EndDate = input.EndDate
	activity.Title = input.Title

	fmt.Println(activity)
	newActivity, err:= s.reporsitory.Save(activity)

	if err != nil {
		return newActivity, err
	}

	return newActivity, nil
}