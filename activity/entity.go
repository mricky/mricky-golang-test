package activity

import "time"


type Activity struct {
	Id 			int
	SkillId	    int
	UserId		int
	Title 		string
	StartDate	string
	EndDate		string
	CreatedAt 	time.Time
	UpdatedAt   time.Time
}
