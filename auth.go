package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateTeacher(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "error in converting string to number. perhaps you send something else than a number as id.",
		})
	}
	// create a token
	token := GenerateToken(ctx.PostForm("password"))
	// create a entry in database
	db.Create(&Teacher{
		FirstName:  ctx.PostForm("first"),
		LastName:   ctx.PostForm("last"),
		PersonalID: id,
		Password:   ctx.PostForm("password"),
		Token:      token,
	})
	// return token
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func CreateStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "error in converting string to number. perhaps you send something else than a number as id.",
		})
	}
  token := GenerateToken(ctx.PostForm("password"))
	db.Create(&Student{
    FirstName:    ctx.PostForm("first"),
    LastName:     ctx.PostForm("last"),
    UniversityID: id,
    Password:     ctx.PostForm("password"),
		Token:        token,
	})
  ctx.JSON(200, gin.H{
		"token": token,
	})
}

func TeacherLogin(ctx *gin.Context)  {
  id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "error in converting string to number. perhaps you send something else than a number as id.",
		})
	}
	var prof Teacher
  result := db.Where(&Teacher{PersonalID: id, Password: ctx.PostForm("password")}).First(&prof)
	if result.Error != nil {
    ctx.JSON(400, gin.H{
			"error": "either username or password is incorrect. recheck the input",
		})
	}
  ctx.JSON(200, gin.H{
		"token": prof.Token,
	})
}

func StudentLogin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
    ctx.JSON(400, gin.H{
			"error": "either username or password is incorrect. recheck the input",
		})
	}
	var student Student
	result := db.Where(&Student{UniversityID: id, Password: ctx.PostForm("password")}).First(&student)
	if result.Error != nil {
    ctx.JSON(400, gin.H{
			"error": "either username or password is incorrect. recheck the input",
		})
	}
  ctx.JSON(200, gin.H{
		"token": student.Token,
	})
}
