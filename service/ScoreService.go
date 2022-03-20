package service

import (
	"github.com/gin-gonic/gin"
	"student_score_sword/models"
	"student_score_sword/repo"
	repoMySqlImpl "student_score_sword/repo/repoMySQLImpl"
)

type ScoreService struct {
	scoreRepo repo.IScoreRepo
}

func (s ScoreService) QueryScore(ctx *gin.Context, id int, id2 int) ([]models.Score, error) {
	score := models.Score{ExamId: id, StudentId: id2}
	query, err := s.scoreRepo.SimpleStructQuery(ctx, score)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func NewScoreService() ScoreService {
	service := ScoreService{scoreRepo: repoMySqlImpl.NewScoreRepoMysql()}
	return service
}
