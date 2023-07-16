package main

import (
	"fmt"
	"log"
	"mricky-golang-test/profile"
	"mricky-golang-test/skill"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:Rahasia123!@tcp(127.0.0.1:8889)/db_mricky_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if (err != nil) {
    	log.Fatal(err.Error())
    }

    fmt.Println("Connection to db is succesfully connected")

	 // scope test repository
	 // Profile
	 profileRepository := profile.ImplProfileRepository(db)
	 data, err := profileRepository.FindAll()
	 data1, err := profileRepository.FindByID(1)

	 fmt.Println(data)
	 fmt.Println(data1)

	 // Skill
	 skillRepository := skill.ImplSkillRepository(db)
	 skills, err := skillRepository.FindAll()
	 skill, err := skillRepository.FindByID(1)

	 fmt.Println(skills)
	 fmt.Println(skill)
}