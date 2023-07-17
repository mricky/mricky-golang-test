# mricky-golang-test

# Steps

1. Create database
2. In main.go please change root and password, please make sure is same with your environment
   dsn := "root:Rahasia123!@tcp(127.0.0.1:8889)/db_mricky_golang?charset=utf8mb4&parseTime=True&loc=Local"
3. Curl Test

- Register

curl --location --request POST 'localhost:8080/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
"name" : "Ricky",
"email" : "mricky.it@gmail.com",
"username" : "mricky.it@gmail.com",
"profile_id": "1",
"password": "password"
}'

- Login
  curl --location --request POST 'localhost:8080/v1/auth/login' \
   --header 'Content-Type: application/json' \
   --data-raw '{
  "email" : "mricky.it@gmail.com",
  "password": "password"
  }'

- Get Skill
  curl --location --request GET 'localhost:8080/v1/skills?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.sAqFOr_3hLrMXrs5oLMtJHqdN_quozlTHAfJegeg02k' \
   --header 'Content-Type: application/json' \
   --data-raw ''

- Get Activitis
  curl --location --request GET 'localhost:8080/v1/activities?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.sAqFOr_3hLrMXrs5oLMtJHqdN_quozlTHAfJegeg02k' \
   --header 'Content-Type: application/json' \
   --data-raw ''

- Save Activities
  curl --location --request POST 'localhost:8080/v1/activities?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.sAqFOr_3hLrMXrs5oLMtJHqdN_quozlTHAfJegeg02k' \
   --header 'Content-Type: application/json' \
   --data-raw '{
  "skill_id" : 2,
  "user_id" : 2,
  "title" : "Bootcame Golang",
  "start_date": "2023-07-25",
  "end_date": "2023-07-27"
  }'
