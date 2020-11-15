package main

import (
	"github.com/gin-gonic/gin"
  "strconv"
)


func CreateCourse(c *gin.Context){
  token := c.PostForm("token")
  var teacher Teacher
  result := db.Where(&Teacher{Token: token}).First(&teacher)
  if result.Error != nil {
    c.JSON(400, gin.H{
      "error": "no such token found. perhaps you are not a teacher",
    })
    return
  }
  code, _:= strconv.Atoi(c.PostForm("code"))
  unit, _:= strconv.Atoi(c.PostForm("unit"))
  capacity, _:= strconv.Atoi(c.PostForm("capacity"))
  result = db.Create(&Course{
    Name: c.PostForm("name"),
    Code: code,
    Unit: unit,
    Capacity: capacity,
    Teacher: teacher,
  })
  if result.Error != nil {
    c.JSON(500, gin.H{
      "error": "cannot add data to database",
    })
    return
  }
  c.JSON(200, gin.H{
    "status": "ok",
  })
}

func SelectCourse(c *gin.Context){
  token := c.PostForm("token")
  var student Student
  result := db.Where(&Student{Token: token}).First(&student)
  if result.Error != nil {
    c.JSON(400, gin.H{
      "error": "no such token found.",
    })
    return
  }
  code, _ := strconv.Atoi(c.PostForm("token"))
  var course Course
  result = db.Where(&Course{Code: code}).First(&course)
  if result.Error != nil {
    c.JSON(400, gin.H{
      "error": "no such course found.",
    })
    return
  }
  if course.Capacity != 0 {
    result = db.Create(&Choice{
      Student: student,
      Course: course,
    })
    course.Capacity--
    db.Save(course)
    c.JSON(200, gin.H{
      "status": "ok",
    })
    return
  }
  c.JSON(400, gin.H{
    "error": "not enough capacity in this course",
  })

}
