# libraries

## GIN

go get -u github.com/gin-gonic/gin

## GORM

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

## Validation

go get github.com/go-playground/validator/v10

## JWT

go get github.com/dgrijalva/jwt-go

# Analisis Data

profiles
id
profile_name
[board, expert, trainer, competitor]

users
id
name
username
password_hash
FK profiles

activities
skill_id (one)
title
description
start_date
end_date

participants
activity_id
user_id

roles = profile
expert can manage activities

Endpoint
activities filterByLogin(User)

# Features

## Auth

v1/auth/login
logout

## User

v1/user

- register
- update
- delete

## Activities

v1/activity

- listBy (skill,user)
- create
- update
- delete

## Skill

v1/skills

- list

# Other

Pagination

## Steps

1. Create Repository and Test x
2. Create Service and Test x
3. Handler Request x
4. Helper and Formater x
5. Test In Posman
6. Unitesting Repostory, Service, Handler
