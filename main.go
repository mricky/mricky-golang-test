package main

import (
	"fmt"
	"log"
	"mricky-golang-test/profile"

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
	 // User
	//  userStruct := user.User{
	// 		Name: "Mohammad Ricky",
	// 		Email: "mricky.it@gmail.com",
	// 		Username: "mricky",
	// 		Password: "rahasia",
	// 		ProfileId: 1,
	//  }

	//  userRepository := user.ImplUserRepository(db)
	//  _,err = userRepository.Save(userStruct)
	//  if(err != nil){
	// 	fmt.Println("Berhasil Insert User")
	//  }
	 // Profile
	 profileRepository := profile.ImplProfileRepository(db)
	 data, err := profileRepository.FindAll()
	 data1, err := profileRepository.FindByID(1)

	 fmt.Println(data)
	 fmt.Println(data1)

}