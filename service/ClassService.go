package service

import (
	"context"
	"student_score_sword/models"
	"student_score_sword/repo"
	repoMySqlImpl "student_score_sword/repo/repoMySQLImpl"
)

type ClassService struct {
	classRepo repo.IClassRepo
}

func (s *ClassService) GetTeacherClass(ctx context.Context, teacherId int) (res []models.Class, err error) {
	res, err = s.classRepo.SimpleStructQuery(ctx, models.Class{TeacherId: teacherId})
	return res, err
}

func NewClassService() ClassService {
	service := ClassService{classRepo: repoMySqlImpl.NewClassRepoMysql()}
	return service
}
