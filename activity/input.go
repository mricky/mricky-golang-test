package activity

type CreateActivityInput struct {
	SkillId 		int 	`json:"skill_id" binding:"required"`
	UserId 			int 	`json:"user_id" binding:"required"`
	Title  			string  `json:"title" binding:"required"`
	StartDate		string  `json:"start_date" binding:"required"`
	EndDate			string  `json:"end_date" binding:"required"`

}
