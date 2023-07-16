CREATE database 'db_mricky_golang'

CREATE TABLE skills (
    id INT NOT NULL AUTO_INCREMENT,
    skill_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
)

CREATE TABLE profiles (
    id INT NOT NULL AUTO_INCREMENT,
    profile_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
)

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    profile_id INT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (profile_id) REFERENCES profiles(id) 
    		ON UPDATE CASCADE
    		ON DELETE CASCADE
);

CREATE TABLE activities (
     id INT NOT NULL AUTO_INCREMENT,
     skill_id int NOT NULL, 
     title VARCHAR(100),
     start_date DATE,
     end_date DATE,
     created_at DATETIME,
     updated_at DATETIME,
     PRIMARY KEY (id),
     FOREIGN KEY (skill_id) REFERENCES skills(id) 
    		ON UPDATE CASCADE
    		ON DELETE CASCADE
);

CREATE TABLE participants
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    activity_id INT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) 
                ON UPDATE CASCADE
                ON DELETE CASCADE,
    FOREIGN KEY (activity_id) REFERENCES activities(id) 
                ON UPDATE CASCADE
                ON DELETE CASCADE
)
