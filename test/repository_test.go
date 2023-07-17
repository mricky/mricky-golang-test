package test

import (
	"fmt"
	"mricky-golang-test/activity"
	"mricky-golang-test/auth"
	"mricky-golang-test/profile"
	"mricky-golang-test/skill"
	"mricky-golang-test/user"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "root:Rahasia123!@tcp(127.0.0.1:8889)/db_mricky_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db
}

func TestProfileFindAll(t *testing.T){
	db := setupTestDB()
	profileRepository := profile.ImplProfileRepository(db)
	profiles, err := profileRepository.FindAll()

	if err != nil {
		panic(err)
	}
	fmt.Println(profiles)
}
func TestProfileFindByID(t *testing.T){
	db := setupTestDB()
	profileRepository := profile.ImplProfileRepository(db)
	profile, err := profileRepository.FindByID(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(profile)
} 
func TestSkillFindAll(t *testing.T){
	db := setupTestDB()
	skillRepository := skill.ImplSkillRepository(db)
	
	skillService := skill.ImplSkillService(skillRepository)
	skills,err := skillService.GetSkills()

	if err != nil {
		panic(err)
	}
	fmt.Println(skills)
}

func TestSkillFindByID(t *testing.T){
	db := setupTestDB()
	skillRepository := skill.ImplSkillRepository(db)
	skill, err := skillRepository.FindByID(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(skill)
} 

func TestGenerateToken(t *testing.T){
	authService := auth.ImplService()

	
	token, err := authService.GenerateToken(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(token)
} 

func TestUserCreate(t *testing.T){
	 db := setupTestDB()
	 
	 userStruct := user.User{
			Name: "Mohammad Ricky",
			Email: "mricky.it@gmail.com",
			Username: "mricky",
			Password: "rahasia",
			ProfileId: 1,
	 }
	
	 userRepository := user.ImplUserRepository(db)
	 
	 _,err := userRepository.Save(userStruct)
	
	 if(err != nil){
		fmt.Println("Berhasil Insert User")
	 }
}

// func TestUserFindAll(t *testing.T){
// 	// next test with handler
// 	db := setupTestDB()
// 	userRepository := user.ImplUserRepository(db)
// 	users, err := userRepository.FindAll()

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(users)
// }


func TestSkillFindByEmail(t *testing.T){
	db := setupTestDB()
	userRepository := user.ImplUserRepository(db)
	userService := user.ImplUserService(userRepository)

	input := user.LoginInput{
		Email: "mricky.it@gmail.com",
		Password: "password",
		
	}
	user, err := userService.Login(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
} 

func TestActivitySave(t *testing.T){
	db := setupTestDB()	 
	activityRepository := activity.ImplActivityRepository(db)

	activity := activity.Activity{
		SkillId: 1,
		UserId: 1,
		Title: "Bootcame Golang",
		StartDate: "2023-06-25",
		EndDate: "2020-06-30",
	}

	_,err := activityRepository.Save(activity)
	
	 if(err != nil){
		fmt.Println("Berhasil Insert Activity")
	 }
}

func TestActivityUpdate(t *testing.T){
	db := setupTestDB()	 
	activityRepository := activity.ImplActivityRepository(db)

	activity := activity.Activity{
		Id: 2,
		SkillId: 1,
		UserId: 1,
		Title: "Bootcame Golang ddd ddd s",
		StartDate: "2023-06-25",
		EndDate: "2020-06-27",
	}

	_,err := activityRepository.Update(activity)
	
	 if(err != nil){
		fmt.Println("Berhasil Insert Activity")
	 }
}

func TestActivityFindAll(t *testing.T){
	// next test with handler
	db := setupTestDB()
	activityRepository := activity.ImplActivityRepository(db)
	
	// Test Service
	activityService := activity.ImplActivityService(activityRepository)
	activities,err := activityService.GetActivities()

	if err != nil {
		panic(err)
	}
	fmt.Println(activities)
}