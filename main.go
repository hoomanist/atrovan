package main

import (
	"github.com/gin-gonic/gin"
  "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	// TODO: switch to postgresql
  db, err = gorm.Open(mysql.Open("hooman:hooman86@tcp(127.0.0.1:3306)/foodly"), &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true ,
  })
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Teacher{})
  db.AutoMigrate(&Student{})
  db.AutoMigrate(&Choice{})
  db.AutoMigrate(&Course{})
	router := gin.Default()
  router.POST("/register/teacher", CreateTeacher)
  router.POST("/register/student", CreateStudent)
  router.POST("/login/teacher", TeacherLogin)
  router.POST("/login/student", StudentLogin)
  router.POST("/submit/course", CreateCourse)
  router.POST("/join/course", SelectCourse)
	router.Run(":3000")
}
