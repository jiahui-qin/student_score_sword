package repoMySqlImpl

import (
	"context"
	"gorm.io/gorm"
	"student_score_sword/models"
	"student_score_sword/repo"
)

type classRepoImpl struct {
	DB *gorm.DB
}

func (c classRepoImpl) SingleInsert(ctx context.Context, score models.Score) (res models.Score, err error) {
	result := c.DB.WithContext(ctx).Create(score)
	if result.Error == nil {
		return score, nil
	}
	return models.Score{}, result.Error
}

func (c classRepoImpl) SimpleStructQuery(ctx context.Context, score models.Score) (res []models.Score, err error) {
	var scoreList []models.Score
	err = c.DB.WithContext(ctx).Debug().Where(&score).Find(&scoreList).Error
	return scoreList, err

}

func NewScoreRepoMysql() repo.IScoreRepo {
	conn, _ := GetConn()
	return &classRepoImpl{conn}
}
