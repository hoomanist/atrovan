package main

import (
  "gorm.io/gorm"
)


type Teacher struct {
  gorm.Model
  FirstName string
  LastName string
  PersonalID int
  Password string
  Token string
}

type Student struct {
  gorm.Model
  FirstName string
  LastName string
  UniversityID int
  Password string
  Token string
}

type Course struct {
  Name string
  Code int
  Unit int
  Capacity int
  Teacher Teacher

}

type Choice struct {
  Student Student
  Course Course
}
