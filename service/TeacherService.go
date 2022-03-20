package service

import (
	"github.com/gin-gonic/gin"
	"student_score_sword/models"
	"student_score_sword/repo"
	repoMySqlImpl "student_score_sword/repo/repoMySQLImpl"
)

type TeacherService struct {
	teacherRepo repo.ITeacherRepo
}

func (s TeacherService) AddTeacher(context *gin.Context, name string) (models.Teacher, error) {
	teacher := models.Teacher{Name: name}
	res, err := s.teacherRepo.SingleInsert(context, teacher)
	return res, err
}

func NewTeacherService() TeacherService {
	service := TeacherService{teacherRepo: repoMySqlImpl.NewTeacherRepoMysql()}
	return service
}
