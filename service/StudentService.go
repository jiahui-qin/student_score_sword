package service

import (
	"github.com/gin-gonic/gin"
	"student_score_sword/models"
	"student_score_sword/repo"
	repoMySqlImpl "student_score_sword/repo/repoMySQLImpl"
)

type StudentService struct {
	studentRepo repo.IStudentRepo
}

func (s StudentService) GetClassStudent(ctx *gin.Context, id int) ([]models.Student, error) {
	student := models.Student{ClassId: id}
	query, err := s.studentRepo.SimpleStructQuery(ctx, student)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func NewStudentService() StudentService {
	service := StudentService{studentRepo: repoMySqlImpl.NewStudentRepoMysql()}
	return service
}
