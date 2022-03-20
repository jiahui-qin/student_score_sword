package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"student_score_sword/models"
	"student_score_sword/service"
)

func main() {

	classService := service.NewClassService()
	examService := service.NewExamService()
	scoreService := service.NewScoreService()
	studentService := service.NewStudentService()
	teacherService := service.NewTeacherService()

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	teacher := r.Group("/teacher")
	{
		teacher.GET("/:id", func(ctx *gin.Context) {
			teacherIdString := ctx.Param("id")
			teacherId, _ := strconv.Atoi(teacherIdString)
			ctx.JSON(200, gin.H{"res": teacherId})
		})
		teacher.POST("/add", func(context *gin.Context) {
			name := context.PostForm("name")
			res, _ := teacherService.AddTeacher(context, name)
			context.JSON(200, gin.H{"res": res})
		})
	}

	exam := r.Group("/exam")
	{
		exam.GET("/query", func(ctx *gin.Context) {
			teacherIdString := ctx.DefaultQuery("teacherId", "0")
			teacherId, _ := strconv.Atoi(teacherIdString)
			classIdString := ctx.Query("classId")
			classId, _ := strconv.Atoi(classIdString)
			res, _ := examService.GetTeacherExam(ctx, teacherId, classId)
			ctx.JSON(200, gin.H{"res": res})
		})
		exam.POST("/add", func(ctx *gin.Context) {
			var exam models.Exam
			ctx.BindJSON(&exam)
			res, _ := examService.AddExam(ctx, exam)
			ctx.JSON(200, gin.H{"res": res})
		})
	}
	class := r.Group("/class")
	{
		class.GET("/:teacherId", func(context *gin.Context) {
			teacherIdString := context.Param("teacherId")
			teacherId, _ := strconv.Atoi(teacherIdString)
			classList, _ := classService.GetTeacherClass(context, teacherId)
			context.JSON(200, gin.H{"res": classList})
		})
		class.POST("/add", func(ctx *gin.Context) {

		})
	}
	student := r.Group("/student")
	{
		student.GET("/query", func(ctx *gin.Context) {
			classIdString := ctx.Query("classId")
			classId, _ := strconv.Atoi(classIdString)
			res, _ := studentService.GetClassStudent(ctx, classId)
			ctx.JSON(200, gin.H{"res": res})
		})
	}
	score := r.Group("/score")
	{
		score.GET("/query", func(ctx *gin.Context) {
			examIdString := ctx.Query("examId")
			examId, _ := strconv.Atoi(examIdString)
			stuIdString := ctx.Query("studentId")
			studentId, _ := strconv.Atoi(stuIdString)
			res, _ := scoreService.QueryScore(ctx, examId, studentId)
			ctx.JSON(200, gin.H{"res": res})
		})
		r.POST("/upload", func(ctx *gin.Context) {
			// Single file
			file, _ := ctx.FormFile("file")
			log.Println(file.Filename)

			// Upload the file to specific dst.
			//ctx.SaveUploadedFile(file, dst)

			ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
