package test

import (
	"fmt"
	"mricky-golang-test/skill"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "root:Rahasia123!@tcp(127.0.0.1:8889)/db_mricky_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db
}

func TestSkillFindAll(t *testing.T){
	// next test with handler
	db := setupTestDB()
	skillRepository := skill.ImplSkillRepository(db)
	skills, err := skillRepository.FindAll()

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