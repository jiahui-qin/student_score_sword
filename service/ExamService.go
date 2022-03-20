package service

import (
	"github.com/gin-gonic/gin"
	"student_score_sword/models"
	"student_score_sword/repo"
	repoMySqlImpl "student_score_sword/repo/repoMySQLImpl"
)

type ExamService struct {
	examRepo repo.IExamRepo
}

func (s ExamService) GetTeacherExam(ctx *gin.Context, teacherId int, classId int) ([]models.Exam, error) {
	exam := models.Exam{TeacherId: teacherId, ClassId: classId}
	query, err := s.examRepo.SimpleStructQuery(ctx, exam)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s ExamService) AddExam(ctx *gin.Context, exam models.Exam) (models.Exam, error) {
	res, err := s.examRepo.SingleInsert(ctx, exam)
	if err != nil {
		return models.Exam{}, err
	}
	return res, nil
}
func NewExamService() ExamService {
	service := ExamService{examRepo: repoMySqlImpl.NewExamRepoMysql()}
	return service
}
