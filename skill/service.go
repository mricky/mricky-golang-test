package skill

type SkillService interface {
	GetSkills()([]Skill, error)
}

type service struct {
	reporsitory SkillRepository
}

func ImplSkillService(reporsitory SkillRepository) *service {
	return &service{reporsitory}
}

func (s *service) GetSkills()([]Skill, error){
	skills, err := s.reporsitory.FindAll()

	if err != nil {
		return skills, err
	}

	return skills, nil
}
